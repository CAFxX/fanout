package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint16; F2 uint8 }
type GenS6 struct { F0 string }
func genFn0(p0 bool) int16 {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: -1}
        _ = fi6
        _ = v7
    }
    v8 := []int{2147483647}
    _ = p0
    _ = v8
    return 155
}
func genFn1() uint8 {
    if false {
        switch 32768 {
        case 32767:
            switch ("" + "a") {
            case "a":
                var v9 uint8 = 250
                _ = v9
            case "hello":
                var v10 float32 = ((float32(0.1) - (0.0 * 0.1)) + 1e10)
                _ = v10
            case "\n\t":
                var v11 uint32 = 0
                _ = v11
            default:
                var v12 int16 = 0
                _ = v12
            }
        default:
            v13 := []int{127}
            _ = v13
        }
    } else {
        switch 1 {
        case -2147483648:
            v14 := GenS2{F0: ""}
            _ = v14
        case -1:
            var v15 float32 = ((0.0 + 3.14159) / 0.1)
            _ = v15
        default:
            v16 := []int{}
            _ = v16
        }
        var v17 uint8 = 128
        var v18 uint8 = (v17 >> ((4294967295) & 63))
        _ = v17
        _ = v18
    }
    if true {
        var v19 uint8 = 1
        v19 += v19
        _ = v19
    } else {
        v20 := []int{128, 4294967295}
        for fi21 := 0; fi21 < 15; fi21++ {
            _ = len(v20)
            var v22 uint8 = 85
            v23 := GenS3{F0: "", F1: -55}
            _ = fi21
            _ = v22
            _ = v23
        }
        switch "\n\t" {
        case "hello":
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 int64 = -9223372036854775808
                v26 := GenS4{F0: 811}
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        _ = v20
    }
    var v27 int64 = -1
    _ = v27
    return 128
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 bool = (!(!(!(!(!((!true) || (1 != 1000)))))))
    var v3 uint16 = 0
    var v4 bool = (4294967295 < 1)
    var v5 float64 = (0.0 / 1e10)
    v28 := GenS4{F0: v1}
    if true {
        v29 := GenS1{F0: -32768}
        v30 := GenS4{F0: (^(v1))}
        v2 = (!v4)
        _ = v29
        _ = v30
    } else {
        if false {
            v31 := GenS3{F0: ("\n\t" + ("hello" + "hello")), F1: 127}
            _ = v31
        } else {
            for fi32 := 0; fi32 < 3; fi32++ {
                switch fi32 {
                case 32767:
                    if (0.0 != 0.0) {
                        v2 = v4
                        v33 := GenS5{F0: (-(fi32)), F1: v3, F2: 0}
                        v34 := GenS6{F0: "foo bar"}
                        _ = v33
                        _ = v34
                    } else {
                        v35 := GenS6{F0: ""}
                        _ = v35
                    }
                }
                v36 := GenS3{F0: "hello", F1: 1}
                for fi37 := 0; fi37 < 8; fi37++ {
                    v38 := []int{-842, fi32}
                    var v39 string = "\n\t"
                    var v40 int = 9223372036854775807
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = fi32
                _ = v36
            }
        }
        if true {
            v41 := []int{9223372036854775807, 32767, -974, 32768}
            _ = v41
        } else {
            fmt.Println(v1)
            v42 := GenS3{F0: "hello", F1: 100}
            _ = v42
        }
        if v4 {
            v43 := []int{4294967295, -162, -9223372036854775808, 1}
            if len(v43) > 0 {
                v43[0] = -756
            }
            _ = v43
        }
    }
    var v44 string = "hello"
    switch v44 {
    case "foo bar":
        var v45 float64 = v5
        _ = v45
    case "":
        fmt.Println(v3, v3, v44)
    case "a":
        v46 := []int{0, -9223372036854775808, -914, 32767}
        _ = v46
    }
    v47 := []int{}
    switch -1 {
    case 1:
        fmt.Println(v5, v44)
    case 9223372036854775807:
        v44 = ("foo bar" + ((v44 + "a") + v44))
    default:
        var v48 uint16 = v3
        _ = v48
    }
    var v49 bool = v2
    if true {
        for fi50 := 0; fi50 < 7; fi50++ {
            var v51 uint64 = 758
            _ = len(v47)
            _ = fi50
            _ = v51
        }
        switch -2147483648 {
        case 42:
            for fi52 := 0; fi52 < 11; fi52++ {
                fmt.Println(fi52)
                var v53 int16 = 1000
                var v54 uint32 = uint32(v1)
                _ = fi52
                _ = v53
                _ = v54
            }
        case 128:
            if true {
                if len(v47) >= 2 {
                    v47 = v47[1:len(v47)]
                }
                var v55 int = -9223372036854775808
                v2 = (v2 || false)
                _ = v55
            }
        case -9223372036854775808:
            fmt.Println(v44)
        }
    } else {
        var v56 int8 = -1
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v49
    _ = v5
}
