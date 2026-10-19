package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint32, p1 uint) int32 {
    if true {
        v5 := []int{}
        var v6 uint16 = 65535
        switch ("foo bar" + "a") {
        case "":
            var v7 uint16 = (v6 | (v6 ^ v6))
            _ = v7
        case "a":
            var v8 string = ("" + "foo bar")
            _ = v8
        default:
            p0 = p0
        }
        _ = v5
        _ = v6
    }
    v9 := GenS1{F0: -1}
    v10 := GenS1{F0: -32768}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int64 = -155
    v2 := []int{128}
    var v3 int16 = -32768
    var v4 string = "\n\t"
    v1 -= (v1 << ((499) & 63))
    if (true && false) {
        v4 = v4
        var v11 int64 = -9223372036854775808
        _ = v11
    }
    v4 = ("a" + "")
    var v12 int8 = -106
    if (128 <= 9223372036854775807) {
        if ((!(v1 != v1)) || false) {
            v12 += v12
            var v13 bool = (v4 != "a")
            v14 := GenS1{F0: (v3 / (((v3 | v3)) | 1))}
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 uint32 = 2147483648
                var v18 uint = 781
                var v19 int32 = -853
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi15
        }
    }
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS2{F0: 2, F1: false, F2: 0}
            v22 := GenS3{F0: 0}
            v1 += 0
            _ = fi20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 uint64 = 2
            _ = fi23
            _ = v24
        }
        if false {
            _ = len(v2)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v12 += v12
        }
    } else {
        if false {
            v4 = ("" + (v4 + ((v4 + (((((v4 + ((v4 + (v4 + v4)) + "foo bar")) + (v4 + (v4 + "foo bar"))) + v4) + "hello") + "\n\t")) + v4)))
        } else {
            var v25 uint16 = 42
            var v26 uint32 = uint32(v25)
            _ = v25
            _ = v26
        }
    }
    v27 := GenS2{F0: 178, F1: (v12 != (v12 >> ((65535) & 63))), F2: 2147483648}
    var v28 float64 = ((0.1 * ((1e10 + float64(1.5)) / 0.1)) * float64(1e10))
    var v29 uint32 = 1
    var v30 int16 = (v3 % ((v3) | 1))
    v2 = append(v2, 4294967295)
    for fi31 := 0; fi31 < 16; fi31++ {
        if true {
            v32 := GenS4{F0: (v1 ^ v1), F1: 32767, F2: 1000}
            _ = v32
        } else {
            v33 := GenS5{F0: true}
            v2 = append(v2, (fi31 | fi31))
            var v34 int8 = 42
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    v2 = append(v2, 9223372036854775807)
    fmt.Println(v3)
    v30 = (v30 / ((1000) | 1))
    fmt.Println(v1, v1, v4)
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
