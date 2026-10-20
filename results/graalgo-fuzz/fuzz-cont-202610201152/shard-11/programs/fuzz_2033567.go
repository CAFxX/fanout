package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 bool; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 bool; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint32) uint32 {
    var v4 int64 = -342
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 bool = false
        var v7 uint = (p0 >> ((65535) & 63))
        if (true && (36 == 147)) {
            var v8 uint32 = (p1 - (^(p1)))
            if v6 {
                var v9 uint8 = 128
                p1 = v8
                _ = v9
            }
            _ = v8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v10 := GenS1{F0: uint64(p1), F1: "\n\t"}
    v11 := GenS1{F0: 381, F1: "a"}
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(p1, fi12)
        switch 127 {
        case 42:
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := []int{fi13}
                v15 := GenS2{F0: false, F1: ((fi13 >> (((^(p0))) & 63)) < (fi13 | -423)), F2: 0}
                p1 -= (^(p1))
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        switch "\n\t" {
        case "a":
            for fi16 := 0; fi16 < 12; fi16++ {
                v4 = (v4 << (((p0 / ((p0) | 1))) & 63))
                var v17 float32 = (((float32(((3.14159 - 0.0) / ((3.14159 / 1.0) + (((-2.25 / 1.0) + (1.5 + 1e10)) + 3.14159)))) + float32(0.0)) / 1e10) / 1.5)
                _ = fi16
                _ = v17
            }
        case "\n\t":
            var v18 bool = (827 > -32768)
            _ = v18
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v4
    return (p1 * (p1 << (((p0 + (p0 * 0))) & 63)))
}

func main() {
    var v1 bool = true
    var v2 int32 = -2147483648
    v3 := make(map[int]string)
    for fi19 := 0; fi19 < 4; fi19++ {
        if v1 {
            fmt.Println(fi19)
            var v20 uint64 = 444
            v21 := GenS3{F0: 1, F1: 4294967295, F2: -32768}
            _ = v20
            _ = v21
        }
        v2 += (-(v2))
        _ = fi19
    }
    v22 := GenS4{F0: 2147483647, F1: 42}
    v23 := []int{}
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 uint64 = 1
        _ = fi24
        _ = v25
    }
    fmt.Println(v1, v2)
    for fi26 := 0; fi26 < 4; fi26++ {
        v1 = v1
        switch (fi26 | -424) {
        case 32768:
            if (!(3.14159 <= 0.0)) {
                var v27 float32 = 3.14159
                var v28 int8 = -128
                var v29 uint16 = 757
                _ = v27
                _ = v28
                _ = v29
            }
        case 128:
            fmt.Println(v1, v2, v1)
        case 978:
            var v30 uint64 = 42
            _ = v30
        }
        v1 = (40 != 2147483648)
        _ = fi26
    }
    if (1 >= -64) {
        for fi31 := 0; fi31 < 1; fi31++ {
            v2 = (v2 | v2)
            v2 -= v2
            _ = fi31
        }
        if true {
            switch ("a" + "a") {
            case "foo bar":
                for fi32 := 0; fi32 < 11; fi32++ {
                    v1 = (0 == 760)
                    _ = fi32
                }
            case "\n\t":
                var v33 int16 = 42
                _ = v33
            case "a":
                fmt.Println(v1, v1)
            }
            v1 = (v1 || v1)
        }
        fmt.Println(v2)
    } else {
        if (642 == 2) {
            var v34 uint16 = 0
            var v35 float32 = 1e10
            if (379 > 0) {
                if (179 != 0) {
                    v2 = 2147483647
                    for fi36 := 0; fi36 < 8; fi36++ {
                        v37 := GenS3{F0: 0, F1: int(v2), F2: 0}
                        _ = fi36
                        _ = v37
                    }
                    if false {
                        var v38 uint32 = 42
                        _ = v38
                    }
                } else {
                    v39 := GenS5{F0: (1 < 42), F1: 127, F2: 22}
                    _ = v39
                }
            } else {
                v1 = v1
            }
            _ = v34
            _ = v35
        } else {
            _ = len(v3)
        }
        var v40 string = ("\n\t" + "")
        v1 = (v1 || false)
        _ = v40
    }
    fmt.Println(v2)
    switch -108 {
    case 9223372036854775807:
        switch ("\n\t" + "a") {
        case "foo bar":
            var v41 float32 = 1e10
            _ = v41
        case "hello":
            var v42 bool = false
            _ = v42
        default:
            var v43 int = 42
            _ = v43
        }
    case -9223372036854775808:
        v2 = v2
    case -1:
        delete(v3, -484)
    }
    if v1 {
        for fi44 := 0; fi44 < 13; fi44++ {
            v45 := GenS6{F0: 1}
            _ = fi44
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
