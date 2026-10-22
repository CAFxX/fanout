package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int8; F2 uint64 }
func genFn0(p0 uint8, p1 int, p2 uint64) int64 {
    for fi4 := 0; fi4 < 5; fi4++ {
        fmt.Println(p0)
        var v5 uint64 = p2
        v6 := GenS1{F0: 2147483648, F1: p2}
        _ = fi4
        _ = v5
        _ = v6
    }
    p2 -= (p2 & uint64(p0))
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]string)
    var v3 string = "hello"
    v3 = ("\n\t" + "")
    fmt.Println(v3, v1)
    for fi7 := 0; fi7 < 1; fi7++ {
        v3 = v3
        v1 -= v1
        v8 := GenS2{F0: 1, F1: 4294967295, F2: v3}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int64 = -1
        if false {
            if false {
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS1{F0: 4294967295, F1: 42}
                    v13 := v2[fi11]
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v14 := GenS1{F0: 244, F1: 4294967295}
                _ = v14
            }
            switch (fi9 % ((fi9) | 1)) {
            case 128:
                _ = len(v2)
            case -9223372036854775808:
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 bool = true
                    var v17 int32 = 42
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            if true {
                for fi18 := 0; fi18 < 9; fi18++ {
                    v1 = v1
                    v3 = "hello"
                    delete(v2, (fi18 ^ fi18))
                    _ = fi18
                }
            }
        } else {
            switch -461 {
            case 42:
                v10 = (-(v10))
            case 0:
                for fi19 := 0; fi19 < 9; fi19++ {
                    v10 = genFn0(0, -725, 82)
                    var v20 int16 = v1
                    var v21 uint8 = 255
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case 2147483647:
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 uint16 = 987
                    _ = fi22
                    _ = v23
                }
            }
            var v24 uint8 = 1
            if false {
                v25 := GenS1{F0: 2147483648, F1: 18446744073709551615}
                v26 := GenS1{F0: 42, F1: 0}
                for fi27 := 0; fi27 < 11; fi27++ {
                    v10 += v10
                    v1 -= v1
                    if (false && false) {
                        var v28 uint = 255
                        _ = v28
                    }
                    _ = fi27
                }
                _ = v25
                _ = v26
            }
            _ = v24
        }
        _ = fi9
        _ = v10
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 uint = 4294967295
        v31 := []int{fi29, 4294967295, (fi29 / ((fi29) | 1)), -9223372036854775808}
        _ = fi29
        _ = v30
        _ = v31
    }
    switch 42 {
    case -2147483648:
        for fi32 := 0; fi32 < 9; fi32++ {
            fmt.Println(v3)
            _ = fi32
        }
    }
    var v33 string = (v3 + ((v3 + "") + v3))
    if (0 != 0) {
        fmt.Println(v1, v3, v3)
        var v34 uint = 255
        _ = v34
    } else {
        for fi35 := 0; fi35 < 8; fi35++ {
            for fi36 := 0; fi36 < 6; fi36++ {
                v3 = ((v3 + "hello") + v33)
                switch (v33 + "foo bar") {
                case "\n\t":
                    v37 := GenS3{F0: true}
                    _ = v37
                case "a":
                    v38 := GenS4{F0: 1, F1: 1, F2: 0}
                    _ = v38
                case "foo bar":
                    var v39 string = "a"
                    _ = v39
                }
                _ = fi36
            }
            var v40 float64 = 3.14159
            _ = fi35
            _ = v40
        }
    }
    var v41 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v41
}
