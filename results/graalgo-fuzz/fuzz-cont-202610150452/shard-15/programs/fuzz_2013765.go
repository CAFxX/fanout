package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 uint64, p1 int32, p2 uint8) uint64 {
    v6 := []int{520, 4294967295, 32768, 712}
    var v7 uint64 = 2
    if true {
        var v8 int64 = 360
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(v8, fi9)
            for fi10 := 0; fi10 < 18; fi10++ {
                if len(v6) > 0 {
                    v6[0] = fi10
                }
                _ = fi10
            }
            _ = len(v6)
            _ = fi9
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -2147483648
    var v3 float64 = (1.5 + 1e10)
    var v4 int32 = v2
    v5 := make(map[uint32]uint32)
    switch ("\n\t" + "foo bar") {
    case "hello":
        switch ("foo bar" + "a") {
        case "foo bar":
            v4 -= 2147483647
        case "hello":
            v11 := GenS1{F0: "foo bar", F1: 2, F2: 498}
            _ = v11
        }
    case "a":
        v12 := GenS2{F0: -1, F1: (109 != -1)}
        _ = v12
    default:
        for fi13 := 0; fi13 < 9; fi13++ {
            if false {
                v14 := GenS2{F0: -119, F1: (3.14159 != -2.25)}
                _ = v14
            }
            var v15 uint64 = 1
            _ = fi13
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        fmt.Println(v4, v3)
        _ = fi16
    }
    if false {
        v17 := []int{-2147483648, 514, 9223372036854775807, 4294967295}
        _ = v17
    } else {
        var v18 int64 = -9223372036854775808
        v1[128] = (-114 > 32767)
        if true {
            var v19 int64 = v18
            _ = v19
        } else {
            v18 = (-(v18))
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint16 = 2
                var v22 string = ""
                v21 = v21
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = v18
    }
    switch -9223372036854775808 {
    case 414:
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS1{F0: ("a" + "hello"), F1: 1, F2: 2147483647}
            var v25 uint32 = 385
            v4 -= 123
            _ = fi23
            _ = v24
            _ = v25
        }
    case -2147483648:
        fmt.Println(v3)
    case -915:
        if false {
            var v26 float64 = float64(3.14159)
            _ = v26
        }
    }
    var v27 int = 0
    fmt.Println(v4)
    var v28 float64 = v3
    fmt.Println(v27, v2)
    v29 := GenS3{F0: 794, F1: 32767}
    if true {
        if true {
            if (true && (255 < 854)) {
                v27 -= 1
            } else {
                for fi30 := 0; fi30 < 11; fi30++ {
                    switch (fi30 ^ v27) {
                    case -1:
                        for fi31 := 0; fi31 < 20; fi31++ {
                            switch ("a" + (((("a" + "\n\t") + ((("foo bar" + (("a" + "\n\t") + "\n\t")) + ("foo bar" + "\n\t")) + "")) + ("\n\t" + "\n\t")) + "foo bar")) {
                            case "foo bar":
                                v32 := GenS2{F0: 42, F1: (2147483647 < (fi31 % ((-848) | 1)))}
                                var v33 uint64 = 0
                                var v34 int64 = 0
                                _ = v32
                                _ = v33
                                _ = v34
                            case "a":
                                var v35 uint64 = 42
                                _ = v35
                            case "\n\t":
                                var v36 int64 = -389
                                _ = v36
                            default:
                                v37 := []int{-1}
                                _ = v37
                            }
                            _ = fi31
                        }
                    case -2147483648:
                        var v38 uint64 = 389
                        _ = v38
                    }
                    _ = fi30
                }
                v39 := []int{v27, v27, v27}
                _ = v39
            }
        } else {
            v40 := GenS2{F0: 55, F1: false}
            _ = v40
        }
        var v41 int32 = 1
        _ = v41
    } else {
        fmt.Println(v28, v3)
    }
    var v42 int8 = 42
    var v43 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
