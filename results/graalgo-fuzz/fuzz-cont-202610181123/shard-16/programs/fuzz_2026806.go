package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int32; F2 uint }
func genFn0() uint {
    v5 := GenS1{F0: 1}
    if false {
        var v6 int = 128
        v6 += (v6 * v6)
        _ = v6
    }
    var v7 uint64 = 1
    var v8 uint32 = 65535
    v9 := GenS2{F0: "\n\t"}
    _ = v5
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := []string{("" + ("hello" + "foo bar")), "hello", "a", ("a" + "foo bar")}
    var v10 bool = (255 < 188)
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint = 0
        _ = fi11
        _ = v12
    }
    v10 = (v10 && (2147483647 <= -2147483648))
    for fi13 := 0; fi13 < 15; fi13++ {
        fmt.Println(fi13, fi13)
        switch "a" {
        case "foo bar":
            var v14 uint64 = 18446744073709551615
            _ = v14
        }
        fmt.Println(v10)
        _ = fi13
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int = (fi15 + (-(fi15)))
        if (!v10) {
            v17 := v2[v16]
            _ = v17
        } else {
            fmt.Println(v10)
        }
        _ = fi15
        _ = v16
    }
    var v18 uint = 18446744073709551615
    v10 = v10
    switch 136 {
    case 9223372036854775807:
        var v19 float64 = 0.1
        _ = v19
    case 257:
        var v20 string = "hello"
        _ = v20
    case -1000:
        fmt.Println(v18, v10)
    }
    if true {
        var v21 int64 = -1
        if v10 {
            for fi22 := 0; fi22 < 5; fi22++ {
                switch (^(fi22)) {
                case 577:
                    v23 := GenS1{F0: 65535}
                    _ = v23
                case 1:
                    fmt.Println(fi22, v21)
                }
                if len(v4) > 0 {
                    v4[0] = ("hello" + (("foo bar" + "a") + "a"))
                }
                _ = fi22
            }
        }
        _ = v21
    } else {
        for fi24 := 0; fi24 < 13; fi24++ {
            v25 := GenS3{F0: true, F1: 308, F2: 235}
            _ = fi24
            _ = v25
        }
        var v26 int8 = 0
        if v10 {
            for fi27 := 0; fi27 < 19; fi27++ {
                if (v10 && true) {
                    switch ("\n\t" + "") {
                    case "foo bar":
                        if len(v4) > 0 {
                            v4[0] = "foo bar"
                        }
                    case "\n\t":
                        var v28 int = (fi27 % ((fi27) | 1))
                        _ = v28
                    }
                }
                delete(v1, 2147483648)
                _ = fi27
            }
            if v10 {
                for fi29 := 0; fi29 < 6; fi29++ {
                    if v10 {
                        v30 := GenS1{F0: 2}
                        _ = v30
                    }
                    _ = fi29
                }
                v18 = v18
            }
            switch 196 {
            case 9223372036854775807:
                var v31 uint = genFn0()
                _ = v31
            case -511:
                var v32 string = ("" + "hello")
                _ = v32
            }
        }
        _ = v26
    }
    v33 := GenS3{F0: (127 == -96), F1: 1, F2: v18}
    _ = len(v4)
    v10 = v10
    v10 = (v18 == (v18 % ((v18) | 1)))
    var v34 int8 = -1
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 int8 = v34
        v36 -= (^(v36))
        fmt.Println(v10, v18)
        _ = fi35
        _ = v36
    }
    v37 := GenS3{F0: (18 >= 128), F1: -512, F2: 0}
    v38 := GenS3{F0: v10, F1: 2147483647, F2: 844}
    if ("\n\t" != "hello") {
        fmt.Println(v10)
        for fi39 := 0; fi39 < 16; fi39++ {
            switch ("foo bar" + "") {
            case "foo bar":
                if v10 {
                    v34 -= (v34 % ((v34) | 1))
                    switch -9223372036854775808 {
                    case 1:
                        var v40 int = fi39
                        _ = v40
                    }
                    switch "" {
                    case "\n\t":
                        v18 = (^(v18))
                    }
                } else {
                    var v41 int64 = -545
                    v34 -= (v34 - (v34 / ((v34) | 1)))
                    var v42 uint32 = 2147483648
                    _ = v41
                    _ = v42
                }
            case "":
                switch (("foo bar" + "a") + ("hello" + (("hello" + "foo bar") + ((("hello" + "foo bar") + "\n\t") + "foo bar")))) {
                case "hello":
                    v18 += v18
                case "a":
                    v43 := GenS3{F0: true, F1: 2147483647, F2: genFn0()}
                    _ = v43
                case "\n\t":
                    var v44 float64 = (0.1 * (float64(-2.25) + ((1e10 + (((3.14159 * (((1e10 + (float64(float32(1.5)) + (1.5 / float64((3.14159 - (3.14159 / -2.25)))))) / 1.0) + 0.0)) * 3.14159) * 1e10)) - 0.0)))
                    _ = v44
                default:
                    if (!v10) {
                        v34 -= -128
                    }
                }
            }
            v10 = (!(-2147483648 < 615))
            for fi45 := 0; fi45 < 1; fi45++ {
                switch (^(fi45)) {
                case 4294967295:
                    var v46 uint = 295
                    _ = v46
                }
                _ = fi45
            }
            _ = fi39
        }
        var v47 float64 = 3.14159
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
}
