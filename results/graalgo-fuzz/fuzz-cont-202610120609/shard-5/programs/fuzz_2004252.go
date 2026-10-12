package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint8 }

func main() {
    var v1 float32 = 3.14159
    v2 := []uint32{420, 994}
    var v3 string = "\n\t"
    v4 := make(map[uint32]int)
    var v5 uint = 2
    if true {
        v1 += v1
        var v6 string = "a"
        _ = v6
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            if len(v2) > 0 {
                v2[0] = 22
            }
            v5 = v5
            for fi8 := 0; fi8 < 2; fi8++ {
                if true {
                    var v9 int16 = -1
                    v9 -= (-(v9))
                    v10 := GenS1{F0: 918}
                    _ = v10
                    _ = v9
                }
                switch (v3 + v3) {
                case "hello":
                    var v11 int32 = 42
                    var v12 uint8 = 0
                    v3 = ((v3 + ("a" + v3)) + (v3 + v3))
                    _ = v11
                    _ = v12
                case "foo bar":
                    v5 = (v5 / (((v5 + (v5 % (((v5 * 42)) | 1)))) | 1))
                case "a":
                    var v13 int = (-(fi7))
                    _ = v13
                }
                _ = fi8
            }
            _ = fi7
        }
        var v14 uint32 = 229
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(v14, v3, v14)
            _ = fi15
        }
        _ = v14
    }
    var v16 uint16 = 467
    if false {
        v1 = v1
        for fi17 := 0; fi17 < 18; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                if (!true) {
                    var v19 string = (("" + v3) + v3)
                    v20 := GenS1{F0: 619}
                    var v21 int8 = 127
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    v22 := GenS2{F0: (v5 << ((2) & 63))}
                    v1 += v1
                    _ = v22
                }
                _ = fi18
            }
            var v23 float32 = 3.14159
            v3 = ("\n\t" + ("a" + (v3 + (v3 + "foo bar"))))
            _ = fi17
            _ = v23
        }
        v1 -= v1
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v24 := GenS1{F0: 1}
    v25 := GenS1{F0: 30}
    v16 = v16
    v1 = 0.0
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 uint16 = 255
        v16 = (^(v16))
        _ = fi26
        _ = v27
    }
    v5 += v5
    if true {
        _ = len(v4)
        v28 := GenS3{F0: v5, F1: 2}
        delete(v4, 42)
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
