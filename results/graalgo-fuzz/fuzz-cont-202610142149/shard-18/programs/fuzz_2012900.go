package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    v1 := []int{32767}
    var v2 bool = false
    v3 := make(map[int]bool)
    v4 := []uint32{688, 65535, 953}
    var v5 int64 = 42
    delete(v3, -911)
    fmt.Println(v2)
    var v6 uint = 255
    fmt.Println(v2)
    v7 := GenS1{F0: 152, F1: 2147483647}
    if v2 {
        var v8 float64 = -2.25
        _ = v8
    }
    var v9 int8 = -82
    if v2 {
        switch ("a" + "") {
        case "a":
            switch "foo bar" {
            case "hello":
                v6 = (v6 << (((v6 ^ v6)) & 63))
            case "foo bar":
                v9 += -44
            case "a":
                if (v2 || (!true)) {
                    v2 = v2
                } else {
                    v9 = v9
                }
            default:
                if ((39 == 42) || (!("hello" > ("" + "hello")))) {
                    for fi10 := 0; fi10 < 6; fi10++ {
                        var v11 string = ("a" + ("a" + ("\n\t" + "")))
                        for fi12 := 0; fi12 < 10; fi12++ {
                            v3[(fi12 * fi12)] = ((!(!("" <= "\n\t"))) && v2)
                            _ = fi12
                        }
                        _ = fi10
                        _ = v11
                    }
                }
            }
        }
    } else {
        v13 := GenS2{F0: 65535}
        _ = v13
    }
    v2 = (-1 <= 0)
    var v14 uint32 = 4294967295
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 string = ((("a" + ("" + "")) + (("" + "a") + ("foo bar" + "foo bar"))) + ("\n\t" + ("foo bar" + ((("\n\t" + "\n\t") + ("foo bar" + ("" + "foo bar"))) + "\n\t"))))
        v17 := GenS2{F0: v14}
        v2 = (int64(fi15) > (v5 | v5))
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: -32768, F1: 513, F2: 4294967295}
    if (false && (!(!false))) {
        _ = len(v3)
        var v19 int64 = (v5 + (v5 / (((v5 ^ v5)) | 1)))
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 string = "foo bar"
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    if v2 {
        var v22 float64 = -2.25
        _ = v22
    }
    fmt.Println(v6, v9)
    var v23 int64 = -628
    if ((!(!v2)) && v2) {
        v4 = append(v4, 603)
        var v24 float32 = 1e10
        _ = v24
    } else {
        if (!false) {
            var v25 int16 = 1000
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 uint8 = 127
                _ = fi26
                _ = v27
            }
            v5 += v5
            _ = v25
        } else {
            if len(v1) > 0 {
                v1[0] = -1
            }
            _ = len(v3)
            var v28 uint = v6
            _ = v28
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
