package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := []int{}
    var v2 float64 = float64(0.1)
    var v3 string = ("\n\t" + ("foo bar" + "foo bar"))
    var v4 float32 = float32((float64(1e10) - (float64((3.14159 / (0.1 * float32((1.5 * float64(1.5)))))) * ((0.0 + 0.1) * (0.1 + v2)))))
    fmt.Println(v3)
    v2 = ((0.0 * (0.1 * v2)) * float64(1e10))
    for fi5 := 0; fi5 < 8; fi5++ {
        switch (fi5 * fi5) {
        case 0:
            fmt.Println(v3, v3)
        case 32767:
            for fi6 := 0; fi6 < 20; fi6++ {
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                var v7 uint64 = 470
                v8 := GenS1{F0: 32767, F1: 42, F2: 0}
                _ = fi6
                _ = v7
                _ = v8
            }
        case 133:
            if (-456 >= 0) {
                switch (fi5 - (-(fi5))) {
                case 42:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -1:
                    if false {
                        for fi9 := 0; fi9 < 18; fi9++ {
                            if (505 == 65535) {
                                _ = len(v1)
                                var v10 bool = (0 <= 256)
                                var v11 int16 = -1
                                _ = v10
                                _ = v11
                            }
                            _ = fi9
                        }
                    }
                case -114:
                    for fi12 := 0; fi12 < 17; fi12++ {
                        v13 := GenS1{F0: -32768, F1: 32767, F2: 739}
                        _ = fi12
                        _ = v13
                    }
                }
            } else {
                v4 -= v4
                v1 = append(v1, (fi5 | (fi5 / (((fi5 & -1)) | 1))))
            }
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS2{F0: 224, F1: 2}
            _ = fi14
            _ = v15
        }
        _ = fi5
    }
    v3 = v3
    v3 = v3
    v3 = v3
    fmt.Println(v4, v4, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
