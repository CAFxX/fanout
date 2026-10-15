package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint8 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
func genFn0(p0 int) string {
    for fi6 := 0; fi6 < 1; fi6++ {
        if (0 <= 4294967295) {
            v7 := GenS1{F0: 2, F1: fi6, F2: 92}
            _ = v7
        }
        var v8 uint64 = 2
        if false {
            var v9 uint32 = 2
            _ = v9
        }
        _ = fi6
        _ = v8
    }
    v10 := []int{-1, p0, p0}
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            if ((fi12 - -1) > p0) {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v13 := GenS1{F0: 65535, F1: (^(p0)), F2: 1}
                _ = v13
            } else {
                fmt.Println(fi11, p0, fi11)
            }
            _ = fi12
        }
        v10 = append(v10, (fi11 & fi11))
        _ = fi11
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        p0 = 127
        for fi15 := 0; fi15 < 9; fi15++ {
            p0 -= fi14
            var v16 float32 = -2.25
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        if len(v10) > 0 {
            v10[0] = -976
        }
        var v18 int = p0
        for fi19 := 0; fi19 < 18; fi19++ {
            if (798 > 42) {
                var v20 float32 = 1e10
                _ = v20
            }
            var v21 uint64 = 648
            _ = fi19
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v10
    return "\n\t"
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 128
    var v3 uint32 = 1
    v4 := make(map[uint32]uint32)
    var v5 int16 = -1
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := []int{517}
            v5 = 297
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            v2 += (v2 >> ((v1) & 63))
            var v26 int64 = 42
            _ = fi25
            _ = v26
        }
        var v27 uint64 = 2
        _ = fi22
        _ = v27
    }
    v3 -= 704
    v5 += 1
    fmt.Println(v5)
    v28 := GenS2{F0: 984, F1: false}
    var v29 int = -96
    fmt.Println(v5)
    _ = len(v4)
    fmt.Println(v3, v2, v29)
    for fi30 := 0; fi30 < 6; fi30++ {
        if false {
            if true {
                var v31 uint8 = 2
                v32 := GenS3{F0: (v3 & (v3 | v3)), F1: v5, F2: 1}
                _ = v31
                _ = v32
            }
            var v33 int16 = v5
            for fi34 := 0; fi34 < 7; fi34++ {
                v3 = 212
                var v35 uint = (^(v1))
                v2 = 163
                _ = fi34
                _ = v35
            }
            _ = v33
        } else {
            var v36 int16 = 545
            v37 := GenS1{F0: 0, F1: -807, F2: v2}
            if false {
                var v38 int32 = int32(v3)
                _ = v38
            } else {
                switch v29 {
                case 2147483647:
                    for fi39 := 0; fi39 < 3; fi39++ {
                        v40 := v4[(v3 | v3)]
                        _ = fi39
                        _ = v40
                    }
                case 32767:
                    var v41 int64 = -1
                    _ = v41
                case 32768:
                    switch "\n\t" {
                    case "\n\t":
                        v36 = v5
                    }
                }
            }
            _ = v36
            _ = v37
        }
        v29 -= (-(fi30))
        _ = fi30
    }
    var v42 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
