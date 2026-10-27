package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint16; F1 int64 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{}
        v9 := GenS1{F0: 2, F1: 32767, F2: ""}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int8 = -92
    fmt.Println(v10)
    v10 = -128
    _ = v10
    return 364
}

func main() {
    var v1 string = "hello"
    var v2 uint64 = 1
    var v3 uint32 = uint32(v2)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    var v6 int = 42
    var v11 bool = (!(v2 < uint64(v3)))
    for fi12 := 0; fi12 < 9; fi12++ {
        fmt.Println(fi12, v1)
        for fi13 := 0; fi13 < 3; fi13++ {
            fmt.Println(v3, v6)
            switch 2147483647 {
            case 1:
                v14 := GenS2{F0: v11}
                _ = v14
            case 9223372036854775807:
                fmt.Println(fi12)
            case 2147483647:
                v15 := GenS2{F0: (v1 < "foo bar")}
                _ = v15
            default:
                if false {
                    var v16 int64 = -9223372036854775808
                    v1 = "a"
                    v17 := GenS1{F0: 255, F1: 42, F2: v1}
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS3{F0: (^(v3)), F1: int16(v2), F2: v11}
                    _ = v18
                }
            }
            if v11 {
                var v19 bool = (!(false || v11))
                var v20 uint32 = v3
                _ = v19
                _ = v20
            } else {
                if false {
                    var v21 string = v1
                    var v22 int64 = -452
                    _ = v21
                    _ = v22
                } else {
                    fmt.Println(v11, v3)
                    v6 = 128
                }
                if (!v11) {
                    var v23 uint32 = (v3 * (v3 - (^(v3))))
                    _ = v23
                } else {
                    v24 := GenS1{F0: 42, F1: 586, F2: ((((((v1 + v1) + (("foo bar" + ("foo bar" + "\n\t")) + "foo bar")) + ("\n\t" + v1)) + ((v1 + "foo bar") + (("a" + (v1 + v1)) + (v1 + (("hello" + "foo bar") + ""))))) + (((v1 + "foo bar") + v1) + v1)) + v1)}
                    _ = v24
                }
            }
            v3 += v3
            v25 := GenS4{F0: -2147483648}
            var v26 uint32 = (^(v3))
            v2 += (v2 - (v2 + 4294967295))
            _ = fi13
            _ = v25
            _ = v26
        }
        _ = fi12
    }
    v27 := GenS5{F0: v3, F1: (v1 + ("\n\t" + v1))}
    var v28 float64 = float64(1.5)
    var v29 uint8 = 255
    var v30 uint32 = (v3 ^ v3)
    delete(v5, (v3 / ((v30) | 1)))
    var v31 int8 = -125
    for fi32 := 0; fi32 < 17; fi32++ {
        if (-1 != -9223372036854775808) {
            v33 := []int{2147483647, (^(fi32)), (fi32 - 0), (v6 + fi32)}
            _ = v33
        }
        _ = fi32
    }
    v34 := v4[("hello" + "foo bar")]
    if v34 {
        if (float64(0.1) == ((v28 - v28) / (0.1 * ((1e10 - (v28 * 1.5)) - (v28 - ((((3.14159 * (1e10 / 1e10)) * 3.14159) / v28) + ((0.1 * v28) / 1e10))))))) {
            var v35 int8 = v31
            _ = v35
        } else {
            delete(v4, "")
        }
    } else {
        switch (v1 + (v1 + "hello")) {
        case "a":
            v11 = (((^(v29)) >= v29) && v34)
        case "":
            for fi36 := 0; fi36 < 4; fi36++ {
                var v37 int = 128
                var v38 int16 = -728
                _ = fi36
                _ = v37
                _ = v38
            }
        case "hello":
            var v39 float64 = float64((0.0 + 0.1))
            _ = v39
        default:
            if (626 < v30) {
                if ((v11 || (false && (42 > 0))) || v11) {
                    v29 = (^(v29))
                    v34 = ("hello" < ("hello" + ""))
                    fmt.Println(v34, v30)
                }
            }
        }
        var v40 uint16 = 72
        _ = v40
    }
    switch "hello" {
    case "\n\t":
        v28 = v28
    case "":
        var v41 int64 = 42
        _ = v41
    case "hello":
        v42 := GenS1{F0: 2, F1: -32768, F2: ((v1 + v1) + ((v1 + "foo bar") + ("a" + ("hello" + ((v1 + ((("foo bar" + "foo bar") + (v1 + (v1 + v1))) + v1)) + "")))))}
        _ = v42
    }
    v43 := GenS4{F0: 0}
    switch v6 {
    case 128:
        switch v1 {
        case "foo bar":
            fmt.Println(v1)
        default:
            switch (v1 + (v1 + "")) {
            case "\n\t":
                if v34 {
                    var v44 bool = v11
                    v3 = 623
                    _ = v44
                }
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    var v45 bool = v34
                    _ = v45
                case "\n\t":
                    v46 := []int{(v6 | (v6 - (v6 & v6))), v6}
                    _ = v46
                }
            case "a":
                v6 = 127
            default:
                v30 = v30
            }
        }
    case 42:
        v30 -= (^(v30))
    default:
        for fi47 := 0; fi47 < 5; fi47++ {
            if (v31 >= (^(v31))) {
                var v48 int32 = 42
                var v49 int16 = 1
                switch v1 {
                case "a":
                    var v50 float32 = 1e10
                    _ = v50
                default:
                    fmt.Println(v34, v31)
                }
                _ = v48
                _ = v49
            }
            v31 = -1
            _ = fi47
        }
    }
    var v51 int8 = v31
    for fi52 := 0; fi52 < 1; fi52++ {
        v53 := GenS6{F0: 32767, F1: -9223372036854775808}
        v54 := []int{v6, (v6 << ((582) & 63))}
        fmt.Println(v2, v11)
        _ = fi52
        _ = v53
        _ = v54
    }
    fmt.Println(v34, v28)
    var v55 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v43
    _ = v5
    _ = v51
    _ = v55
    _ = v6
}
