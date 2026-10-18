package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{1, 0, 1}
    v3 := make(map[string]uint32)
    if (962 < 32767) {
        var v4 uint = 0
        for fi5 := 0; fi5 < 14; fi5++ {
            v4 -= (v4 << ((2) & 63))
            var v6 int64 = 0
            v4 -= 255
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        _ = len(v2)
        _ = fi7
    }
    var v8 string = ("hello" + "\n\t")
    v8 = ((((v8 + (("\n\t" + ("a" + "")) + (v8 + v8))) + v8) + v8) + "")
    var v9 int32 = -743
    if false {
        for fi10 := 0; fi10 < 3; fi10++ {
            switch (v8 + (v8 + v8)) {
            case "":
                if (0 != 900) {
                    for fi11 := 0; fi11 < 5; fi11++ {
                        fmt.Println(fi11, v8)
                        fmt.Println(fi10, fi11)
                        var v12 float32 = 3.14159
                        _ = fi11
                        _ = v12
                    }
                }
            case "a":
                v2 = append(v2, 191)
            case "hello":
                var v13 uint = 884
                _ = v13
            default:
                if true {
                    v9 = (v9 - -917)
                    v9 += 2147483647
                } else {
                    var v14 int64 = int64(fi10)
                    fmt.Println(v8)
                    v15 := GenS1{F0: 0, F1: (v14 + v14), F2: 128}
                    _ = v14
                    _ = v15
                }
            }
            _ = fi10
        }
        var v16 float64 = (-2.25 * (-2.25 * (0.1 / (0.0 - (0.1 / 0.1)))))
        _ = v16
    }
    v8 = ((((v8 + (((((v8 + (v8 + v8)) + v8) + ("\n\t" + ("hello" + (v8 + ("\n\t" + "hello"))))) + v8) + ((("" + ("hello" + v8)) + v8) + ((((v8 + v8) + v8) + ((v8 + v8) + ((v8 + "") + v8))) + "foo bar")))) + ((v8 + v8) + v8)) + "a") + (v8 + ((("foo bar" + "hello") + v8) + "\n\t")))
    for fi17 := 0; fi17 < 12; fi17++ {
        v8 = "a"
        for fi18 := 0; fi18 < 14; fi18++ {
            fmt.Println(fi17, fi17)
            v19 := GenS1{F0: 19, F1: -985, F2: 1}
            switch -932 {
            case -1:
                _ = len(v2)
            case 291:
                fmt.Println(v8, fi17, fi17)
            case 127:
                v20 := GenS2{F0: 65535}
                _ = v20
            }
            var v21 int64 = 9223372036854775807
            delete(v3, ("a" + v8))
            var v22 int32 = v9
            _ = fi18
            _ = v19
            _ = v21
            _ = v22
        }
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        _ = fi17
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
