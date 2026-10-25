package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint8; F1 int64 }
func genFn0() int64 {
    v6 := GenS1{F0: 1}
    var v7 uint = 255
    var v8 int32 = 42
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint8 = 42
    var v2 string = ("hello" + "a")
    var v3 int = 32768
    var v4 int8 = 0
    v5 := make(map[int]int)
    v3 = (v3 * 0)
    fmt.Println(v1)
    for fi9 := 0; fi9 < 17; fi9++ {
        switch -9223372036854775808 {
        case 2147483647:
            switch v2 {
            case "\n\t":
                if ((!((false && (180 >= -1)) || false)) || true) {
                    v10 := []int{805}
                    var v11 int8 = v4
                    _ = v10
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        v13 := GenS1{F0: 1}
                        v14 := v5[fi9]
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
            default:
                if false {
                    var v15 uint = 4294967295
                    _ = v15
                } else {
                    delete(v5, -659)
                    v16 := GenS1{F0: 561}
                    _ = v16
                }
            }
        default:
            v17 := GenS2{F0: v3, F1: 331}
            _ = v17
        }
        switch v2 {
        case "foo bar":
            v18 := GenS2{F0: v3, F1: 479}
            _ = v18
        case "a":
            switch v2 {
            case "\n\t":
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(v3, v2)
                    _ = fi19
                }
            default:
                v2 = ""
            }
        default:
            var v20 int16 = 1
            _ = v20
        }
        switch (fi9 ^ fi9) {
        case 127:
            switch v2 {
            case "\n\t":
                var v21 int32 = 2147483647
                _ = v21
            case "foo bar":
                var v22 uint = 4294967295
                _ = v22
            default:
                var v23 uint16 = 255
                _ = v23
            }
        case 0:
            var v24 uint64 = 42
            _ = v24
        default:
            v4 = v4
        }
        _ = fi9
    }
    v2 = v2
    if true {
        switch (v3 + 642) {
        case 32768:
            v25 := GenS3{F0: true, F1: (-(v3)), F2: uint32(v1)}
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 uint32 = 0
            var v28 uint32 = (v27 % ((v27) | 1))
            _ = fi26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 uint = 2
            _ = fi29
            _ = v30
        }
    }
    var v31 string = ((v2 + v2) + ("\n\t" + (v2 + ("\n\t" + v2))))
    for fi32 := 0; fi32 < 7; fi32++ {
        for fi33 := 0; fi33 < 9; fi33++ {
            v31 = ("foo bar" + v2)
            fmt.Println(v1, v2, v1)
            _ = fi33
        }
        _ = fi32
    }
    var v34 uint = 2
    v35 := GenS1{F0: 65535}
    switch v3 {
    case 4294967295:
        for fi36 := 0; fi36 < 11; fi36++ {
            switch (v3 | (v3 & v3)) {
            case -194:
                v37 := GenS2{F0: fi36, F1: 42}
                _ = v37
            case 2147483647:
                v34 = 0
            default:
                if true {
                    fmt.Println(v1, fi36)
                } else {
                    v4 += 42
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v39 := GenS4{F0: v4, F1: (v1 * (^(v1))), F2: 1000}
                        if (("hello" + v31) >= v2) {
                            var v40 int8 = v4
                            v1 -= v1
                            v40 = v4
                            _ = v40
                        } else {
                            var v41 uint16 = 741
                            _ = v41
                        }
                        v42 := GenS4{F0: v4, F1: (v1 & v1), F2: 0}
                        _ = fi38
                        _ = v39
                        _ = v42
                    }
                }
            }
            switch v31 {
            case "a":
                _ = len(v5)
            case "hello":
                var v43 string = (v2 + "\n\t")
                _ = v43
            default:
                v44 := []int{-2147483648}
                _ = v44
            }
            if (251 >= 745) {
                if ((!((^(v34)) != v34)) || (genFn0() <= 9223372036854775807)) {
                    v1 += (^(v1))
                    v1 -= 2
                    switch -1 {
                    case 128:
                        v31 = "\n\t"
                    case 654:
                        switch v31 {
                        case "foo bar":
                            v45 := GenS4{F0: v4, F1: (v1 & v1), F2: -1}
                            _ = v45
                        case "\n\t":
                            v46 := GenS1{F0: 120}
                            _ = v46
                        }
                    }
                } else {
                    var v47 uint8 = v1
                    switch v2 {
                    case "a":
                        v48 := GenS5{F0: 21, F1: int64(v34)}
                        _ = v48
                    default:
                        v49 := v5[v3]
                        _ = v49
                    }
                    _ = v47
                }
            }
            _ = fi36
        }
    default:
        var v50 bool = (false || true)
        _ = v50
    }
    v2 = v31
    switch 9223372036854775807 {
    case -1:
        v4 = v4
    case -943:
        _ = len(v5)
    default:
        v2 = v31
    }
    fmt.Println(v31, v2)
    if true {
        var v51 float32 = 1.5
        _ = v51
    }
    fmt.Println(v2, v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
