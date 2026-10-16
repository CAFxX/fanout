package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 bool }
type GenS2 struct { F0 uint32; F1 uint8; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 string; F2 bool }
func genFn0(p0 int64, p1 uint8) uint64 {
    var v4 string = "a"
    var v5 uint8 = (p1 | p1)
    p1 = (v5 >> ((300) & 63))
    v5 = (v5 % ((v5) | 1))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint64 = 215
    var v3 bool = true
    var v6 int32 = -679
    for fi7 := 0; fi7 < 5; fi7++ {
        if (255 <= 18446744073709551615) {
            var v8 int = fi7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 string = (("hello" + "hello") + ("hello" + "\n\t"))
                fmt.Println(v2, v10)
                _ = fi9
                _ = v10
            }
            v3 = (v2 == v2)
            if (v3 && v3) {
                var v11 uint8 = 224
                var v12 bool = (1 != 32767)
                _ = v11
                _ = v12
            } else {
                var v13 int32 = v6
                _ = v13
            }
        }
        v14 := []int{2147483647, 32767}
        _ = fi7
        _ = v14
    }
    if (!v3) {
        switch "foo bar" {
        case "a":
            v15 := []int{453, -1}
            _ = v15
        }
    } else {
        var v16 int32 = v6
        fmt.Println(v16, v16, v6)
        switch 889 {
        case 127:
            v17 := GenS1{F0: 42, F1: v3, F2: ("\n\t" != "")}
            _ = v17
        default:
            var v18 int64 = 1
            _ = v18
        }
        _ = v16
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 int16 = 132
        v2 = 1
        fmt.Println(v2, v2)
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 32767, F1: false, F2: (9223372036854775807 > 353)}
    if true {
        for fi22 := 0; fi22 < 10; fi22++ {
            fmt.Println(v2, v3)
            v23 := []int{(fi22 | fi22), 0, fi22, fi22}
            _ = fi22
            _ = v23
        }
        v24 := []int{1, 1, 32768, -697}
        v24 = append(v24, 2147483647)
        _ = v24
    } else {
        var v25 string = ((("\n\t" + "\n\t") + "a") + ("\n\t" + "a"))
        v26 := GenS2{F0: 2147483648, F1: 0, F2: v3}
        v3 = true
        _ = v25
        _ = v26
    }
    fmt.Println(v3, v3)
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 string = ""
        v29 := GenS3{F0: 42}
        v28 = v28
        _ = fi27
        _ = v28
        _ = v29
    }
    if (("" + "\n\t") > "") {
        var v30 bool = v3
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 uint64 = v2
            v33 := []int{}
            v34 := GenS3{F0: 42}
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        var v35 float64 = 0.1
        _ = v30
        _ = v35
    }
    v1 = (v1 - (^(v1)))
    for fi36 := 0; fi36 < 16; fi36++ {
        v6 = 42
        if v3 {
            var v37 int32 = -1
            _ = v37
        }
        switch ("a" + "\n\t") {
        case "a":
            switch ("a" + ("\n\t" + ("foo bar" + "\n\t"))) {
            case "\n\t":
                if (false || v3) {
                    v38 := GenS3{F0: 42}
                    var v39 int64 = int64(v6)
                    fmt.Println(v1, v2, v3)
                    _ = v38
                    _ = v39
                }
            }
        case "hello":
            v40 := GenS1{F0: -33, F1: v3, F2: (false && (v2 != (^(v2))))}
            _ = v40
        }
        _ = fi36
    }
    fmt.Println(v3, v6)
    for fi41 := 0; fi41 < 18; fi41++ {
        v42 := GenS4{F0: (v6 % ((v6) | 1)), F1: ("\n\t" + ("" + "hello")), F2: v3}
        v43 := []int{}
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v6
}
