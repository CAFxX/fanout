package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
func genFn0(p0 uint32) uint16 {
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int64 = 42
            _ = fi8
            _ = v9
        }
        if true {
            v10 := GenS1{F0: ("foo bar" < "\n\t")}
            switch "a" {
            case "\n\t":
                p0 -= p0
            case "":
                v11 := GenS1{F0: true}
                _ = v11
            case "a":
                p0 += p0
            default:
                for fi12 := 0; fi12 < 19; fi12++ {
                    p0 -= p0
                    p0 += (p0 + (p0 << ((1) & 63)))
                    v13 := []int{fi7}
                    _ = fi12
                    _ = v13
                }
            }
            switch ((("foo bar" + (("\n\t" + ("hello" + ("hello" + "\n\t"))) + "hello")) + "\n\t") + "foo bar") {
            case "\n\t":
                var v14 int = fi7
                _ = v14
            }
            _ = v10
        }
        _ = fi7
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int16 = -1
        if true {
            if true {
                fmt.Println(v16, v16, fi15)
                if (-1 != -932) {
                    fmt.Println(fi15)
                } else {
                    var v17 uint8 = 127
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v16 = v16
                        _ = fi18
                    }
                    v19 := []int{(fi15 - (fi15 ^ (fi15 / ((-453) | 1)))), (fi15 / ((fi15) | 1)), (fi15 - fi15), 127}
                    _ = v17
                    _ = v19
                }
            }
            v20 := GenS1{F0: true}
            _ = v20
        } else {
            v21 := []int{(fi15 + 4294967295)}
            if (true && (false || (((1e10 - (3.14159 / ((3.14159 / 1e10) / -2.25))) >= 1.5) && false))) {
                _ = len(v21)
                var v22 string = "foo bar"
                _ = v22
            }
            _ = v21
        }
        switch "hello" {
        case "foo bar":
            v23 := []int{fi15}
            _ = v23
        case "hello":
            v24 := GenS1{F0: true}
            _ = v24
        case "a":
            v16 = (v16 * (v16 - 42))
        default:
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := []int{-2147483648}
                var v27 int = (fi25 % (((fi25 * fi15)) | 1))
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        _ = fi15
        _ = v16
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := []int{}
        if (true || (3.14159 > 0.1)) {
            v30 := GenS2{F0: (1000 == 42), F1: 42, F2: -1}
            _ = v30
        } else {
            var v31 uint64 = 34
            fmt.Println(p0, p0, fi28)
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    v32 := GenS3{F0: 128, F1: p0}
    fmt.Println(p0)
    _ = p0
    _ = v32
    return 65535
}

func main() {
    var v1 float64 = 0.1
    var v2 string = "hello"
    var v3 string = v2
    var v4 uint64 = 1
    var v5 uint = 0
    var v6 string = v3
    fmt.Println(v2, v5)
    if (v3 == v3) {
        var v33 int16 = -1
        if (-217 < 42) {
            if false {
                var v34 float32 = (3.14159 / ((-2.25 - 1.5) - 1.5))
                v35 := []int{9223372036854775807}
                var v36 string = ("a" + v2)
                _ = v34
                _ = v35
                _ = v36
            } else {
                var v37 int = 9223372036854775807
                _ = v37
            }
            v5 -= v5
            var v38 string = v6
            _ = v38
        }
        _ = v33
    } else {
        if false {
            v4 = (v4 & v4)
        }
    }
    fmt.Println(v2)
    for fi39 := 0; fi39 < 2; fi39++ {
        v40 := []int{fi39, fi39}
        if false {
            var v41 uint = (v5 | v5)
            _ = v41
        }
        _ = fi39
        _ = v40
    }
    fmt.Println(v6)
    fmt.Println(v3, v3)
    v42 := []int{1, -1, 32767, 0}
    var v43 uint16 = 42
    if false {
        fmt.Println(v5, v5, v2)
        v1 = v1
    }
    v44 := GenS1{F0: (-1 != 0)}
    for fi45 := 0; fi45 < 12; fi45++ {
        v3 = (v2 + v6)
        v1 -= v1
        switch (fi45 >> ((v5) & 63)) {
        case -9223372036854775808:
            v46 := GenS4{F0: (v5 + (^(v5)))}
            _ = v46
        case 9223372036854775807:
            v47 := GenS2{F0: true, F1: (v5 ^ 42), F2: -1}
            _ = v47
        }
        _ = fi45
    }
    if (true || (!(!(!false)))) {
        var v48 uint64 = (v4 | (^(v4)))
        _ = v48
    } else {
        var v49 uint8 = 136
        _ = v49
    }
    v50 := GenS5{F0: (v5 + (v5 * v5))}
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v50
    _ = v6
}
