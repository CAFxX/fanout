package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint8; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
func genFn0() int {
    var v7 int32 = 1
    var v8 int64 = -9223372036854775808
    var v9 int8 = -1
    v8 = (v8 << ((255) & 63))
    v7 = v7
    _ = v7
    _ = v8
    _ = v9
    return 32768
}
func genFn1(p0 int16) string {
    v10 := GenS1{F0: 0, F1: 42, F2: 2}
    switch -1 {
    case -9223372036854775808:
        p0 = (p0 % ((42) | 1))
    default:
        v11 := []int{int(p0)}
        _ = v11
    }
    v12 := []int{1}
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint16 = 0
        v15 := GenS2{F0: 4294967295}
        if (!(2147483647 == 1)) {
            v16 := GenS2{F0: 255}
            _ = v16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v10
    _ = v12
    return ("hello" + "foo bar")
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]int)
    var v3 uint32 = 2
    var v4 int16 = int16(v1)
    v5 := make(map[string]uint32)
    var v6 uint8 = 255
    if false {
        if true {
            for fi17 := 0; fi17 < 6; fi17++ {
                v4 += 0
                _ = fi17
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                v19 := []int{}
                switch "foo bar" {
                case "foo bar":
                    var v20 int = genFn0()
                    _ = v20
                }
                v21 := GenS1{F0: 557, F1: 65535, F2: uint16(v3)}
                _ = fi18
                _ = v19
                _ = v21
            }
            var v22 int32 = -551
            _ = v22
        }
        v23 := []int{32768, 4294967295}
        _ = v23
    } else {
        var v24 uint8 = (v6 * (^(v6)))
        v25 := GenS1{F0: 18446744073709551615, F1: 65535, F2: 541}
        _ = v24
        _ = v25
    }
    var v26 bool = true
    switch ("hello" + ("hello" + "")) {
    case "\n\t":
        var v27 bool = ((v3 + v3) != v3)
        _ = v27
    case "":
        v28 := GenS1{F0: 42, F1: uint(v3), F2: 696}
        _ = v28
    }
    v6 -= v6
    for fi29 := 0; fi29 < 1; fi29++ {
        v26 = v26
        if (!((42 > (-(v1))) || v26)) {
            for fi30 := 0; fi30 < 10; fi30++ {
                var v31 int8 = 0
                v31 = v31
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 17; fi32++ {
                v6 -= v6
                var v33 string = "foo bar"
                _ = fi32
                _ = v33
            }
            v34 := []int{127}
            _ = v34
        } else {
            if (!true) {
                v3 -= 2147483648
                var v35 int32 = v1
                _ = v35
            } else {
                v6 -= (v6 * v6)
            }
            switch fi29 {
            case 32768:
                var v36 int = genFn0()
                _ = v36
            case 883:
                var v37 float32 = 1e10
                _ = v37
            case -636:
                switch (fi29 + 32767) {
                case -813:
                    fmt.Println(fi29)
                }
            }
        }
        switch "a" {
        case "\n\t":
            var v38 int16 = int16(v3)
            _ = v38
        case "a":
            var v39 int64 = 1
            _ = v39
        default:
            var v40 uint64 = 18446744073709551615
            _ = v40
        }
        _ = fi29
    }
    var v41 string = "foo bar"
    v42 := GenS3{F0: (v4 << ((255) & 63)), F1: 128, F2: 983}
    if (v6 < v6) {
        var v43 int8 = 17
        var v44 int64 = 9223372036854775807
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        for fi46 := 0; fi46 < 1; fi46++ {
            if ((uint64(v3) != 1) || true) {
                if true {
                    var v47 bool = false
                    _ = v47
                } else {
                    v48 := GenS2{F0: 65535}
                    v4 = v4
                    _ = v48
                }
            } else {
                v49 := []int{}
                v26 = (v1 < (^(v1)))
                var v50 string = v41
                _ = v49
                _ = v50
            }
            _ = fi46
        }
        v51 := []int{(fi45 >> ((255) & 63))}
        _ = fi45
        _ = v51
    }
    for fi52 := 0; fi52 < 9; fi52++ {
        var v53 int16 = (v4 ^ v4)
        _ = fi52
        _ = v53
    }
    switch v41 {
    case "foo bar":
        fmt.Println(v4)
    case "\n\t":
        var v54 int = 296
        _ = v54
    case "hello":
        v55 := GenS4{F0: 42}
        _ = v55
    default:
        fmt.Println(v3, v26, v4)
    }
    v3 -= (^(v3))
    v56 := GenS5{F0: 822}
    v1 = (v1 + v1)
    v57 := GenS2{F0: 1}
    for fi58 := 0; fi58 < 7; fi58++ {
        if (!((3.14159 != 1e10) || false)) {
            var v59 float64 = 0.0
            _ = v59
        } else {
            v26 = v26
            fmt.Println(v1, v4, v6)
        }
        _ = fi58
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
