package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32 }
func genFn0() int {
    v7 := GenS1{F0: 42}
    switch "\n\t" {
    case "foo bar":
        var v8 int64 = -9223372036854775808
        _ = v8
    case "":
        var v9 uint16 = 0
        _ = v9
    case "a":
        if (false || (0 == 1)) {
            if true {
                var v10 uint8 = 127
                if (49 < 867) {
                    var v11 string = (("a" + "") + "hello")
                    for fi12 := 0; fi12 < 5; fi12++ {
                        var v13 uint64 = 4294967295
                        v14 := []int{}
                        var v15 int16 = 32767
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    _ = v11
                } else {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        var v17 bool = true
                        v18 := []int{}
                        var v19 int32 = 2147483647
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    v20 := []int{}
                    _ = v20
                }
                _ = v10
            }
            v21 := []int{-183, -2147483648}
            var v22 float64 = -2.25
            _ = v21
            _ = v22
        }
    }
    v23 := []int{9223372036854775807, 9223372036854775807}
    v24 := GenS2{F0: false}
    _ = v23
    _ = v24
    _ = v7
    return 42
}
func genFn1(p0 int16) string {
    var v25 int16 = p0
    p0 -= v25
    v26 := []int{}
    var v27 int = 150
    _ = p0
    _ = v25
    _ = v26
    _ = v27
    return "hello"
}

func main() {
    var v1 float32 = ((((3.14159 * 1e10) - 1.5) + float32(0.0)) / (0.1 + ((0.0 * (-2.25 / 1e10)) + 1e10)))
    var v2 int64 = 0
    v3 := make(map[string]bool)
    var v4 int32 = -15
    var v5 float32 = (v1 + v1)
    var v6 int32 = v4
    for fi28 := 0; fi28 < 2; fi28++ {
        fmt.Println(v5, v1)
        v29 := GenS2{F0: (65535 < 65535)}
        switch (("\n\t" + ("\n\t" + "a")) + ("\n\t" + ("foo bar" + ("hello" + "")))) {
        case "hello":
            for fi30 := 0; fi30 < 10; fi30++ {
                v31 := v3["a"]
                _ = fi30
                _ = v31
            }
        case "a":
            v32 := GenS2{F0: (((255 > 2) && (2147483648 >= 2)) || true)}
            _ = v32
        default:
            switch (fi28 / ((fi28) | 1)) {
            case -2147483648:
                switch -1 {
                case 895:
                    for fi33 := 0; fi33 < 16; fi33++ {
                        v34 := []int{fi33, fi28, (fi28 % ((128) | 1)), (^(fi28))}
                        v34 = append(v34, (fi33 | 1))
                        _ = fi33
                        _ = v34
                    }
                case 1:
                    for fi35 := 0; fi35 < 8; fi35++ {
                        fmt.Println(fi28, v1, fi28)
                        var v36 uint8 = 127
                        _ = fi35
                        _ = v36
                    }
                }
            }
        }
        _ = fi28
        _ = v29
    }
    fmt.Println(v5, v4)
    var v37 string = "\n\t"
    v38 := GenS3{F0: 94, F1: 42}
    fmt.Println(v37, v37, v5)
    v39 := GenS4{F0: 0}
    v40 := GenS1{F0: -32768}
    var v41 int32 = v6
    v5 -= -2.25
    v6 += (^(v41))
    fmt.Println(v5, v4, v6)
    if ((v2 % (((v2 * 9223372036854775807)) | 1)) != (v2 + 1)) {
        var v42 int = 933
        _ = v42
    }
    v1 -= (-2.25 / (v1 - v5))
    var v43 uint64 = 1
    fmt.Println(v1)
    switch 42 {
    case 0:
        v44 := GenS5{F0: 932}
        _ = v44
    case 103:
        for fi45 := 0; fi45 < 8; fi45++ {
            for fi46 := 0; fi46 < 9; fi46++ {
                var v47 int = 0
                switch (v47 << ((18446744073709551615) & 63)) {
                case 9223372036854775807:
                    v2 = v2
                }
                v48 := GenS3{F0: 42, F1: 65535}
                _ = fi46
                _ = v47
                _ = v48
            }
            if true {
                for fi49 := 0; fi49 < 3; fi49++ {
                    v6 = v41
                    fmt.Println(v37, v41, v5)
                    _ = fi49
                }
                for fi50 := 0; fi50 < 4; fi50++ {
                    var v51 uint16 = 42
                    v6 += (v41 * v4)
                    v52 := GenS3{F0: int8(v43), F1: 138}
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            }
            v6 -= 0
            _ = fi45
        }
    default:
        for fi53 := 0; fi53 < 16; fi53++ {
            v54 := GenS5{F0: 359}
            _ = fi53
            _ = v54
        }
    }
    var v55 uint64 = (v43 % (((v43 / (((v43 % (((v43 << ((448) & 63))) | 1))) | 1))) | 1))
    if (42 < 383) {
        var v56 float32 = -2.25
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v5
    _ = v55
    _ = v6
}
