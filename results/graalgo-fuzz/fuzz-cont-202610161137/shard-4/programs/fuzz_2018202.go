package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
func genFn0() int {
    switch 169 {
    case 42:
        var v4 uint8 = 82
        _ = v4
    case 4294967295:
        var v5 uint16 = 65535
        _ = v5
    default:
        var v6 float64 = 0.1
        _ = v6
    }
    var v7 bool = (159 >= 255)
    v7 = (109 <= 0)
    _ = v7
    return -2147483648
}
func genFn1(p0 uint32) int32 {
    switch ("foo bar" + ("\n\t" + (("" + "") + "foo bar"))) {
    case "a":
        if (42 == 2147483647) {
            v8 := []int{99, 127}
            var v9 string = "\n\t"
            _ = v8
            _ = v9
        }
    case "\n\t":
        p0 = (p0 / ((693) | 1))
    }
    switch ("foo bar" + "") {
    case "hello":
        if true {
            v10 := []int{127}
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if (!(true || (2 != 0))) {
            v12 := GenS1{F0: 1, F1: 0, F2: -1}
            _ = v12
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            p0 = p0
            switch fi11 {
            case 127:
                p0 = 65535
            default:
                v14 := []int{}
                _ = v14
            }
            _ = fi13
        }
        var v15 int32 = -3
        _ = fi11
        _ = v15
    }
    var v16 uint8 = 239
    if true {
        v17 := []int{476, 32768, -9223372036854775808, -2147483648}
        v18 := GenS1{F0: 33, F1: (v16 ^ (^(v16))), F2: 42}
        for fi19 := 0; fi19 < 13; fi19++ {
            if true {
                v20 := GenS1{F0: uint(v16), F1: (v16 & 42), F2: 42}
                fmt.Println(p0)
                var v21 bool = true
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                switch 1 {
                case 32767:
                    fmt.Println(p0)
                case 9223372036854775807:
                    fmt.Println(v16)
                }
                if true {
                    v16 = 2
                } else {
                    v16 -= v16
                }
                var v23 uint64 = 0
                _ = fi22
                _ = v23
            }
            var v24 int64 = 9223372036854775807
            _ = fi19
            _ = v24
        }
        _ = v17
        _ = v18
    } else {
        v25 := GenS1{F0: 2, F1: 255, F2: 100}
        _ = v25
    }
    _ = p0
    _ = v16
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    var v3 uint8 = 75
    v26 := []int{9223372036854775807}
    var v27 float64 = 0.0
    var v28 uint64 = 445
    switch 127 {
    case 9223372036854775807:
        v29 := GenS1{F0: 42, F1: v3, F2: 29}
        _ = v29
    default:
        delete(v1, ("a" + "\n\t"))
    }
    var v30 string = "foo bar"
    var v31 uint32 = 42
    for fi32 := 0; fi32 < 13; fi32++ {
        fmt.Println(v30)
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 uint = 0
            _ = fi33
            _ = v34
        }
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
        _ = fi32
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        if (v31 > v31) {
            fmt.Println(v31, v27)
            v27 -= 1.5
        }
        v36 := GenS1{F0: 0, F1: uint8(v31), F2: 1}
        switch fi35 {
        case -165:
            var v37 bool = false
            _ = v37
        }
        _ = fi35
        _ = v36
    }
    v38 := GenS1{F0: 39, F1: v3, F2: 127}
    v27 = v27
    if (0.0 >= 1e10) {
        for fi39 := 0; fi39 < 20; fi39++ {
            v40 := GenS1{F0: 42, F1: v3, F2: 0}
            v26 = append(v26, (fi39 >> ((396) & 63)))
            _ = fi39
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 4; fi41++ {
            var v42 int8 = 51
            if ((-1 < 42) && (!(v30 >= v30))) {
                for fi43 := 0; fi43 < 12; fi43++ {
                    switch "hello" {
                    case "":
                        v26 = append(v26, (^(fi43)))
                        v3 -= 1
                    case "hello":
                        var v44 uint16 = 2
                        _ = v44
                    }
                    fmt.Println(fi43, fi41)
                    _ = fi43
                }
            } else {
                if false {
                    for fi45 := 0; fi45 < 11; fi45++ {
                        var v46 int = fi41
                        _ = fi45
                        _ = v46
                    }
                } else {
                    for fi47 := 0; fi47 < 3; fi47++ {
                        v27 += 0.0
                        _ = fi47
                    }
                }
                v30 = v30
                var v48 float64 = v27
                _ = v48
            }
            v31 -= v31
            _ = fi41
            _ = v42
        }
        var v49 int64 = 42
        _ = v49
    }
    var v50 int64 = 1
    v51 := GenS1{F0: 65535, F1: (v3 << ((0) & 63)), F2: 42}
    fmt.Println(v50, v27, v31)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v50
    _ = v51
}
