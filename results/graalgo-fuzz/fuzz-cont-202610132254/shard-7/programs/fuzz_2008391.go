package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
func genFn0(p0 uint16, p1 uint32, p2 int) int16 {
    v7 := GenS1{F0: 218, F1: -2147483648}
    var v8 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 88
    var v3 uint = 1
    var v4 uint8 = (v2 & v2)
    var v5 int16 = -821
    var v6 uint = (v3 * v3)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint16 = 255
        _ = fi9
        _ = v10
    }
    v4 += (v4 / (((^(v2))) | 1))
    if true {
        v11 := GenS2{F0: 354}
        fmt.Println(v5, v3)
        _ = v11
    }
    var v12 bool = true
    v13 := []int{-1, -901, int(v3)}
    var v14 int8 = 100
    switch "a" {
    case "":
        var v15 uint = 329
        _ = v15
    default:
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS3{F0: "hello"}
            switch ("\n\t" + ("foo bar" + (("foo bar" + "\n\t") + ""))) {
            case "":
                var v18 float64 = 3.14159
                _ = v18
            case "foo bar":
                fmt.Println(v3, v2)
            }
            for fi19 := 0; fi19 < 4; fi19++ {
                fmt.Println(fi16, v12, v4)
                var v20 string = ("\n\t" + ("a" + "foo bar"))
                if (true && ((float32((-2.25 * 3.14159)) * 1e10) >= 0.1)) {
                    var v21 uint16 = 140
                    var v22 int8 = -128
                    _ = v21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
    }
    switch "" {
    case "":
        for fi23 := 0; fi23 < 2; fi23++ {
            v6 = v6
            _ = fi23
        }
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 16; fi25++ {
            if v12 {
                v12 = false
                var v26 int64 = -1
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = v26
            } else {
                var v27 float32 = 3.14159
                _ = v27
            }
            if len(v13) > 0 {
                v13[0] = fi24
            }
            _ = fi25
        }
        switch fi24 {
        case 78:
            _ = len(v1)
        case 446:
            fmt.Println(v4, v14)
        default:
            var v28 uint64 = 18446744073709551615
            _ = v28
        }
        _ = fi24
    }
    var v29 uint64 = 0
    if ((v29 % ((v29) | 1)) >= v29) {
        switch "foo bar" {
        case "\n\t":
            var v30 uint16 = uint16(v4)
            _ = v30
        default:
            v4 -= (^(v2))
        }
        v6 = (v3 >> ((v3) & 63))
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        fmt.Println(v12, v5)
        var v32 bool = ((v29 >= (v29 | v29)) && true)
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 uint16 = 255
            _ = fi33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    switch 946 {
    case 2147483647:
        switch -784 {
        case -2147483648:
            v6 += (^(v6))
        case 575:
            fmt.Println(v6, v14)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
