package packer

import (
	"github.com/HamedFazaely/gopos/iso"
	"github.com/HamedFazaely/gopos/iso/ierrors"
	"github.com/HamedFazaely/gopos/iso/logger"
	"github.com/HamedFazaely/gopos/iso/util"
)

//ISOMessagePacker packs the outer most message
type ISOMessagePacker struct {
	packerBase
	bm iso.Bitmap
}

//NewISOMessagePacker is a factory
func NewISOMessagePacker(lgr logger.Logger, bitmap iso.Bitmap) *ISOMessagePacker {
	return &ISOMessagePacker{
		packerBase: packerBase{
			Logger: lgr,
		},
		bm: bitmap,
	}
}

//Pack packs the composite
func (im *ISOMessagePacker) Pack(c iso.Component) ([]byte, error) {

	if !c.IsComposite() {
		im.Log(logger.Error, ierrors.ErrPackerNotCompatible.Error())
		return nil, ierrors.ErrPackerNotCompatible
	}
	result, err := im.GetPackedMTI(c)
	if err != nil {
		im.logError(err)
		return nil, err
	}

	bm, err := im.calculateBitmap(c, im.bm)
	if err != nil {
		im.logError(err)
		return nil, err
	}
	result = append(result, bm...)
	sf, err := im.sortAndPack(c)
	if err != nil {
		im.logError(err)
		return nil, err
	}
	for _, p := range sf {
		result = append(result, p...)
	}
	return result, nil
}

func (im *ISOMessagePacker) calculateBitmap(c iso.Component, bm iso.Bitmap) ([]byte, error) {
	return util.CalculateBitmap(c, bm)
}

func (im *ISOMessagePacker) sortAndPack(c iso.Component) ([][]byte, error) {
	return util.SortAndPackFields(c)
}
