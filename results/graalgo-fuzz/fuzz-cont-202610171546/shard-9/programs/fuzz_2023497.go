package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint8; F2 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16; F1 uint16 }
func genFn0() uint {
    v7 := GenS1{F0: 2147483647, F1: 42}
    switch 9223372036854775807 {
    case -2147483648:
        var v8 uint = 255
        _ = v8
    default:
        var v9 string = ("hello" + "\n\t")
        _ = v9
    }
    var v10 int = 0
    if (true && (true && (0 >= -1))) {
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 4; fi12++ {
                if (2147483647 > 2147483647) {
                    v13 := GenS2{F0: 0}
                    fmt.Println(v10, fi12, fi11)
                    v14 := GenS3{F0: 65535, F1: 2}
                    _ = v13
                    _ = v14
                } else {
                    var v15 int = fi12
                    v16 := GenS4{F0: 1000, F1: 42, F2: -32768}
                    v15 = 207
                    _ = v15
                    _ = v16
                }
                v10 = 4294967295
                _ = fi12
            }
            v17 := GenS5{F0: 518}
            _ = fi11
            _ = v17
        }
    }
    v18 := []int{-1, 0}
    _ = v10
    _ = v18
    _ = v7
    return 255
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int64 = 1
    v3 := []int{}
    v4 := make(map[string]uint32)
    var v5 int16 = -181
    var v6 float32 = (0.1 * (float32(((3.14159 + 3.14159) - -2.25)) * 1.5))
    var v19 uint = 65535
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS6{F0: false}
        var v22 uint64 = 606
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int8 = 42
    switch "hello" {
    case "foo bar":
        _ = len(v3)
    default:
        if true {
            if true {
                switch -2147483648 {
                case -9223372036854775808:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 32768:
                    v19 += genFn0()
                }
            } else {
                fmt.Println(v1)
            }
            v24 := GenS7{F0: 42, F1: 0}
            _ = v24
        } else {
            v25 := GenS1{F0: -2147483648, F1: 0}
            fmt.Println(v19)
            _ = v25
        }
    }
    if true {
        for fi26 := 0; fi26 < 10; fi26++ {
            for fi27 := 0; fi27 < 14; fi27++ {
                if (!true) {
                    var v28 uint = 123
                    _ = v28
                } else {
                    fmt.Println(v19, v23)
                }
                var v29 int64 = v2
                _ = fi27
                _ = v29
            }
            var v30 uint32 = 65535
            _ = fi26
            _ = v30
        }
        switch ("" + "\n\t") {
        case "a":
            switch 9223372036854775807 {
            case -1:
                for fi31 := 0; fi31 < 6; fi31++ {
                    var v32 int16 = v5
                    if ((!true) && (492 != v5)) {
                        v32 -= (v32 | v5)
                        v19 = v19
                        v1 = v1
                    } else {
                        var v33 bool = (((v23 << ((v19) & 63)) == (-(v23))) && (!true))
                        switch fi31 {
                        case 128:
                            fmt.Println(v33, v32)
                        case 0:
                            v34 := GenS4{F0: -1, F1: 128, F2: v5}
                            _ = v34
                        case -2147483648:
                            v32 -= v5
                        }
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
            case 9223372036854775807:
                fmt.Println(v1, v2)
            }
        default:
            var v35 float64 = 3.14159
            _ = v35
        }
        var v36 int32 = 2147483647
        _ = v36
    }
    fmt.Println(v23)
    for fi37 := 0; fi37 < 12; fi37++ {
        switch (fi37 / (((fi37 - (fi37 * (fi37 - (fi37 % ((int(v1)) | 1)))))) | 1)) {
        case 42:
            var v38 uint8 = 127
            _ = v38
        case -704:
            var v39 int16 = 1
            _ = v39
        case 327:
            var v40 int64 = v2
            _ = v40
        }
        var v41 int8 = v23
        _ = fi37
        _ = v41
    }
    v1 = v1
    if len(v3) > 0 {
        v3[0] = 837
    }
    v1 -= 2
    switch ("hello" + "foo bar") {
    case "\n\t":
        var v42 uint8 = 128
        _ = v42
    case "foo bar":
        _ = len(v3)
    case "":
        v43 := GenS5{F0: (-(v2))}
        _ = v43
    default:
        if (255 < 42) {
            _ = len(v3)
            delete(v4, "")
        } else {
            v4["hello"] = 2147483648
            v44 := v4["a"]
            fmt.Println(v5, v44)
            _ = v44
        }
    }
    v45 := v4[""]
    for fi46 := 0; fi46 < 2; fi46++ {
        var v47 int64 = 394
        if len(v3) > 0 {
            v3[0] = -718
        }
        var v48 int16 = (v5 % ((v5) | 1))
        _ = fi46
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
