package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 4294967295
    var v3 uint8 = 109
    var v4 uint = 131
    for fi5 := 0; fi5 < 20; fi5++ {
        if ((!(false && true)) || true) {
            v4 = uint(v1)
            for fi6 := 0; fi6 < 3; fi6++ {
                v2 -= uint64(v4)
                _ = fi6
            }
            var v7 uint = 588
            _ = v7
        }
        switch "" {
        case "a":
            var v8 int32 = -1
            _ = v8
        default:
            v9 := GenS1{F0: (v1 - 4294967295)}
            _ = v9
        }
        fmt.Println(v3, v1, v3)
        _ = fi5
    }
    v2 = (v2 % ((42) | 1))
    if true {
        var v10 float32 = 3.14159
        var v11 int32 = 0
        _ = v10
        _ = v11
    }
    var v12 float32 = (1e10 - ((((0.1 - 1e10) / ((3.14159 * -2.25) / 1e10)) + 1e10) - -2.25))
    fmt.Println(v12, v3)
    var v13 int16 = int16(v3)
    var v14 float64 = ((0.1 * 0.1) - (1.5 - -2.25))
    for fi15 := 0; fi15 < 20; fi15++ {
        switch "hello" {
        case "a":
            var v16 float64 = (0.1 + v14)
            _ = v16
        case "":
            if ((v3 >> (((v4 % (((v4 | v4)) | 1))) & 63)) <= v3) {
                for fi17 := 0; fi17 < 8; fi17++ {
                    v14 = (v14 / v14)
                    fmt.Println(fi15, v13, v4)
                    var v18 uint64 = 1
                    _ = fi17
                    _ = v18
                }
            } else {
                switch ("\n\t" + (("a" + "foo bar") + "\n\t")) {
                case "\n\t":
                    v13 = 32767
                case "foo bar":
                    for fi19 := 0; fi19 < 17; fi19++ {
                        var v20 bool = false
                        _ = fi19
                        _ = v20
                    }
                }
            }
        case "hello":
            var v21 string = ("a" + ("hello" + ("" + "a")))
            _ = v21
        }
        v4 += (^(v4))
        v3 -= v3
        _ = fi15
    }
    v14 += v14
    var v22 string = ""
    v12 -= v12
    v23 := []int{1, -155, 32768, int(v1)}
    var v24 uint8 = (v3 % (((v3 ^ v3)) | 1))
    fmt.Println(v1, v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
