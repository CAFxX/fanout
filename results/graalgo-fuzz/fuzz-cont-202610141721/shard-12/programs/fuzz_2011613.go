package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int; F1 uint16; F2 uint }
type GenS6 struct { F0 string }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{(fi7 / (((fi6 ^ (^(fi6)))) | 1)), (fi6 ^ fi6)}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v9 uint8 = 95
            _ = fi7
            _ = v8
            _ = v9
        }
        if (false || true) {
            v10 := []int{}
            fmt.Println(fi6)
            _ = v10
        } else {
            var v11 uint8 = 128
            fmt.Println(v11, v11)
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 uint64 = 668
                var v14 int64 = -9223372036854775808
                v11 += 0
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := []int{-99}
            fmt.Println(fi6, fi6)
            fmt.Println(fi6)
            _ = fi15
            _ = v16
        }
        _ = fi6
    }
    var v17 int = -9223372036854775808
    if ((1.5 > (float64(1.5) + 1.5)) && false) {
        if (3.14159 == 3.14159) {
            var v18 uint64 = 4294967295
            _ = v18
        } else {
            v19 := []int{}
            var v20 uint = 868
            var v21 int32 = 0
            _ = v19
            _ = v20
            _ = v21
        }
    } else {
        v22 := GenS1{F0: 596}
        v23 := GenS2{F0: 2147483648, F1: 1}
        switch "foo bar" {
        case "a":
            var v24 string = "a"
            _ = v24
        default:
            v25 := GenS3{F0: 384, F1: -32}
            _ = v25
        }
        _ = v22
        _ = v23
    }
    var v26 uint = 4294967295
    _ = v17
    _ = v26
    return 2
}

func main() {
    v1 := []int{711, 9223372036854775807, -9223372036854775808, 4294967295}
    var v2 float32 = 0.1
    var v3 string = "hello"
    var v4 int64 = -1
    v5 := make(map[int]int)
    var v27 uint64 = 132
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := GenS2{F0: 792, F1: -10}
        if (v27 > (^(v27))) {
            _ = len(v1)
        } else {
            v30 := GenS1{F0: v27}
            var v31 float32 = -2.25
            _ = v30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    switch 469 {
    case 32768:
        if (("a" + "\n\t") > "hello") {
            v2 = (((((v2 + (v2 / 1.5)) * v2) - float32(0.0)) / v2) + (v2 + float32(((1e10 - (0.1 * (1e10 - -2.25))) * (0.0 * 3.14159)))))
            var v32 uint = 997
            _ = v32
        } else {
            v3 = "a"
            var v33 uint32 = 84
            v34 := v5[-947]
            _ = v33
            _ = v34
        }
    case 128:
        v3 = (("" + ((v3 + "") + v3)) + "a")
    default:
        var v35 uint = 742
        _ = v35
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        v3 = "\n\t"
        _ = fi36
    }
    switch v3 {
    case "foo bar":
        var v37 float32 = v2
        _ = v37
    case "":
        switch ((v3 + "a") + (v3 + v3)) {
        case "foo bar":
            var v38 string = "\n\t"
            _ = v38
        default:
            switch v3 {
            case "\n\t":
                if (2 > 18446744073709551615) {
                    var v39 uint = 164
                    _ = v39
                } else {
                    v40 := GenS4{F0: v3, F1: genFn0()}
                    var v41 int16 = -886
                    _ = v40
                    _ = v41
                }
            default:
                var v42 int8 = 42
                _ = v42
            }
        }
    case "hello":
        v27 += (v27 * v27)
    }
    var v43 int8 = -1
    switch v3 {
    case "a":
        switch -9223372036854775808 {
        case -204:
            fmt.Println(v4)
        case 127:
            var v44 uint8 = 2
            _ = v44
        case -9223372036854775808:
            v45 := GenS3{F0: 42, F1: 2147483647}
            _ = v45
        }
    case "\n\t":
        var v46 bool = true
        _ = v46
    }
    var v47 int8 = v43
    _ = len(v1)
    v27 = (v27 << ((25) & 63))
    v48 := GenS2{F0: 65535, F1: v47}
    v43 = v47
    v4 -= v4
    v5[32768] = 128
    v3 = v3
    switch 1 {
    case 676:
        var v49 uint = 723
        _ = v49
    case 4294967295:
        var v50 bool = (true && (0 == 65535))
        _ = v50
    case 32768:
        for fi51 := 0; fi51 < 10; fi51++ {
            var v52 int8 = v43
            if (1000 <= -14) {
                var v53 uint64 = (^(v27))
                v54 := GenS5{F0: fi51, F1: 32767, F2: 255}
                _ = v53
                _ = v54
            } else {
                delete(v5, (fi51 >> ((18446744073709551615) & 63)))
                v27 -= v27
            }
            _ = fi51
            _ = v52
        }
    default:
        v43 += 22
    }
    if (true && false) {
        v43 = (v43 % ((v43) | 1))
        v55 := GenS6{F0: (v3 + ("a" + (v3 + v3)))}
        _ = v55
    } else {
        v27 = v27
        switch (v3 + v3) {
        case "\n\t":
            v56 := GenS2{F0: 821, F1: 1}
            _ = v56
        case "foo bar":
            for fi57 := 0; fi57 < 7; fi57++ {
                v58 := GenS1{F0: 4294967295}
                _ = fi57
                _ = v58
            }
        case "":
            fmt.Println(v3, v4)
        }
        v3 = "a"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v48
    _ = v5
}
