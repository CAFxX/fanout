package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
func genFn0(p0 bool, p1 int8) uint8 {
    if true {
        var v7 int = -1
        fmt.Println(v7, v7, p0)
        v8 := GenS1{F0: 201, F1: (v7 + 128), F2: (v7 <= (^(v7)))}
        _ = v7
        _ = v8
    } else {
        var v9 int8 = p1
        _ = v9
    }
    var v10 uint32 = 1
    var v11 uint8 = 127
    v12 := []int{127, 127, 127}
    var v13 uint8 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return (v13 ^ (v11 / ((uint8(v10)) | 1)))
}
func genFn1(p0 uint32, p1 int) bool {
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int64 = -9223372036854775808
        if (255 != 65535) {
            v16 := GenS1{F0: 1000, F1: 32768, F2: true}
            var v17 uint64 = uint64(p0)
            _ = v16
            _ = v17
        } else {
            var v18 bool = ("\n\t" >= "hello")
            _ = v18
        }
        v19 := GenS1{F0: -1, F1: -2147483648, F2: false}
        _ = fi14
        _ = v15
        _ = v19
    }
    var v20 uint = 383
    if ((3.14159 - 0.0) >= -2.25) {
        switch (("\n\t" + "hello") + ("a" + "a")) {
        case "":
            v21 := []int{}
            _ = v21
        case "hello":
            fmt.Println(p1, p1)
        case "a":
            var v22 int64 = int64(p1)
            _ = v22
        }
        p1 = 9223372036854775807
    }
    if false {
        var v23 bool = false
        fmt.Println(p0, p1, v20)
        p0 += p0
        _ = v23
    } else {
        v20 = (v20 * (^(v20)))
    }
    if false {
        for fi24 := 0; fi24 < 12; fi24++ {
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 int8 = 33
                _ = fi25
                _ = v26
            }
            fmt.Println(fi24, fi24, fi24)
            v27 := GenS1{F0: -32768, F1: 32768, F2: false}
            _ = fi24
            _ = v27
        }
    } else {
        p0 = (p0 % ((42) | 1))
    }
    _ = p0
    _ = p1
    _ = v20
    return (1 == 128)
}

func main() {
    var v1 uint8 = 255
    var v2 string = ("hello" + "hello")
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    v6 := []int{127, -301}
    var v28 uint8 = (v1 + 42)
    v28 += (^(v28))
    var v29 uint8 = (v28 >> ((42) & 63))
    switch ((v2 + ("hello" + v2)) + v2) {
    case "a":
        v2 = v2
    case "\n\t":
        v30 := GenS1{F0: 576, F1: 973, F2: (!false)}
        _ = v30
    case "":
        v1 += (^(v28))
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        for fi32 := 0; fi32 < 12; fi32++ {
            v1 += v29
            switch (((v2 + v2) + ("foo bar" + ("a" + v2))) + v2) {
            case "\n\t":
                v28 += v1
            default:
                var v33 float32 = 3.14159
                _ = v33
            }
            _ = fi32
        }
        v34 := GenS1{F0: -1, F1: (fi31 - (^(fi31))), F2: (!false)}
        for fi35 := 0; fi35 < 11; fi35++ {
            var v36 bool = (fi35 != (fi35 - 1))
            var v37 int8 = 42
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi31
        _ = v34
    }
    if (true || (-89 <= 1)) {
        v29 += 112
        v29 -= (v1 % (((v29 | (v28 | v28))) | 1))
    } else {
        for fi38 := 0; fi38 < 12; fi38++ {
            _ = len(v6)
            _ = fi38
        }
    }
    fmt.Println(v2)
    var v39 float64 = 0.1
    fmt.Println(v2, v1, v2)
    v1 -= (v29 - 42)
    if true {
        if true {
            var v40 int16 = 1000
            _ = v40
        }
    } else {
        switch 85 {
        case -524:
            var v41 uint8 = (v28 ^ (v1 / ((v29) | 1)))
            _ = v41
        case 32768:
            var v42 bool = (-919 > -1)
            _ = v42
        case 435:
            _ = len(v3)
        default:
            var v43 int = -445
            _ = v43
        }
        var v44 int8 = -1
        _ = v44
    }
    v45 := GenS2{F0: 1}
    switch "foo bar" {
    case "hello":
        var v46 uint64 = 18446744073709551615
        _ = v46
    case "":
        switch "\n\t" {
        case "foo bar":
            v1 -= (v29 - (v1 + (v29 % ((v28) | 1))))
        case "\n\t":
            var v47 int = 42
            _ = v47
        default:
            if (1 != 1000) {
                if (-1 < 1) {
                    v2 = ((v2 + (("\n\t" + "foo bar") + v2)) + v2)
                    fmt.Println(v1, v28)
                    v48 := GenS3{F0: -1}
                    _ = v48
                }
                v49 := v3[4294967295]
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v49
            }
        }
    default:
        delete(v3, 32767)
    }
    v29 += (v1 >> ((0) & 63))
    switch 1 {
    case 1:
        v50 := GenS4{F0: 0}
        _ = v50
    case 128:
        var v51 float64 = 1.5
        _ = v51
    }
    if false {
        if len(v6) > 0 {
            v6[0] = 2147483647
        }
        v4["a"] = (true && (!(!(42 > 32768))))
        v28 += v1
    }
    v52 := GenS2{F0: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v45
    _ = v5
    _ = v52
    _ = v6
}
