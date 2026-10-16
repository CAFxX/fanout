package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint32; F2 int16 }
type GenS3 struct { F0 int8; F1 uint64; F2 int }
type GenS4 struct { F0 int8 }

func main() {
    var v1 string = ("foo bar" + ("\n\t" + "foo bar"))
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 int32 = 927
    var v5 int32 = v4
    var v6 uint = 255
    var v7 bool = true
    v8 := []int{72}
    var v9 uint32 = 642
    var v10 uint8 = 154
    switch v1 {
    case "foo bar":
        for fi11 := 0; fi11 < 11; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: -566}
                v7 = v7
                if ((v10 - 255) >= (v10 + 175)) {
                    _ = len(v8)
                    _ = len(v2)
                }
                var v14 bool = (!((v4 << ((v6) & 63)) <= -2147483648))
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    }
    if false {
        v15 := GenS1{F0: 9223372036854775807}
        fmt.Println(v4, v4, v6)
        _ = v15
    }
    var v16 uint = (v6 >> (((v6 >> (((v6 ^ 0)) & 63))) & 63))
    for fi17 := 0; fi17 < 11; fi17++ {
        if true {
            var v18 uint32 = v9
            if true {
                for fi19 := 0; fi19 < 17; fi19++ {
                    var v20 uint8 = 42
                    v7 = v7
                    switch fi17 {
                    case 131:
                        _ = len(v3)
                    case 32767:
                        var v21 int16 = 172
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
                var v22 float64 = 3.14159
                v7 = ((!v7) || v7)
                _ = v22
            }
            _ = v18
        } else {
            for fi23 := 0; fi23 < 16; fi23++ {
                v7 = true
                _ = fi23
            }
            if ((v9 + v9) == 74) {
                if false {
                    var v24 int = (^(fi17))
                    v6 = v6
                    for fi25 := 0; fi25 < 1; fi25++ {
                        v26 := GenS2{F0: 0, F1: (v9 + v9), F2: 915}
                        fmt.Println(v1, v6, fi17)
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                }
                v1 = v1
                for fi27 := 0; fi27 < 6; fi27++ {
                    v8 = append(v8, int(v16))
                    v28 := GenS3{F0: 1, F1: 2, F2: (fi27 - (fi17 << (((v6 & v6)) & 63)))}
                    _ = fi27
                    _ = v28
                }
            } else {
                for fi29 := 0; fi29 < 3; fi29++ {
                    if (!(!false)) {
                        v10 = (^(v10))
                    } else {
                        var v30 string = v1
                        v30 = v1
                        _ = v30
                    }
                    delete(v3, (fi29 + (fi17 * 128)))
                    v31 := GenS4{F0: 0}
                    _ = fi29
                    _ = v31
                }
            }
        }
        var v32 int16 = -537
        _ = fi17
        _ = v32
    }
    var v33 int16 = 609
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
