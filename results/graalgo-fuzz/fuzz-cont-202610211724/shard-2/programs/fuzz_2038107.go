package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint64 }

func main() {
    var v1 int = 1
    v2 := make(map[string]bool)
    var v3 uint16 = 966
    v4 := make(map[string]string)
    var v5 uint32 = 2
    var v6 int16 = 1000
    var v7 bool = (1 > 1)
    v1 += -302
    var v8 int = (-(v1))
    v8 = (v1 ^ (v8 / ((9223372036854775807) | 1)))
    if (2147483647 <= 596) {
        var v9 uint16 = (^(v3))
        if v7 {
            v10 := []int{(v1 & v1), (v1 & v1), 301, (^(v1))}
            _ = v10
        }
        _ = v9
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            v7 = v7
            v12 := v4["a"]
            v13 := GenS1{F0: 0}
            _ = fi11
            _ = v12
            _ = v13
        }
        v4[""] = ("\n\t" + ("foo bar" + "a"))
        v5 = (v5 - (v5 / ((v5) | 1)))
    }
    v14 := GenS1{F0: 18446744073709551615}
    var v15 int64 = 42
    for fi16 := 0; fi16 < 5; fi16++ {
        v4[""] = ("hello" + "")
        if (!v7) {
            for fi17 := 0; fi17 < 3; fi17++ {
                if (!(2 <= 2)) {
                    fmt.Println(fi16, v5, fi17)
                } else {
                    var v18 int8 = 127
                    _ = v18
                }
                v8 += (-(v8))
                var v19 float32 = -2.25
                _ = fi17
                _ = v19
            }
            v3 = v3
            if true {
                fmt.Println(fi16, v6)
                v3 = uint16(v5)
            } else {
                var v20 float64 = 1.5
                _ = v20
            }
        } else {
            var v21 int16 = 1
            _ = v21
        }
        for fi22 := 0; fi22 < 2; fi22++ {
            switch (fi22 - -233) {
            case -907:
                v23 := GenS1{F0: 65535}
                if (v7 || false) {
                    switch ("a" + "hello") {
                    case "foo bar":
                        switch "foo bar" {
                        case "\n\t":
                            var v24 float32 = ((1.5 - -2.25) + 1e10)
                            _ = v24
                        default:
                            var v25 uint8 = 2
                            _ = v25
                        }
                    default:
                        fmt.Println(v8)
                    }
                }
                if true {
                    v7 = true
                    v7 = (4294967295 >= 4294967295)
                } else {
                    if ("a" <= "foo bar") {
                        var v26 int = fi16
                        v3 += uint16(v5)
                        if false {
                            var v27 float32 = 3.14159
                            var v28 int32 = -2147483648
                            v15 = v15
                            if v7 {
                                var v29 uint = 4294967295
                                var v30 uint64 = 0
                                _ = v29
                                _ = v30
                            }
                            fmt.Println(fi22)
                            v31 := GenS1{F0: 853}
                            v32 := []int{fi16, (v8 + int(v5))}
                            var v33 uint = 1
                            _ = v27
                            _ = v28
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        _ = v26
                    } else {
                        v3 += v3
                        switch "hello" {
                        case "":
                            v34 := GenS2{F0: (v5 ^ v5)}
                            _ = v34
                        case "hello":
                            v35 := GenS2{F0: v5}
                            _ = v35
                        }
                        v36 := GenS3{F0: -1}
                        _ = v36
                    }
                    if v7 {
                        fmt.Println(v1)
                    }
                    if (v3 > (v3 | v3)) {
                        var v37 float64 = 3.14159
                        if (!true) {
                            var v38 float32 = (0.0 - 1e10)
                            var v39 uint8 = 128
                            v1 = (v8 + (^(fi22)))
                            _ = v38
                            _ = v39
                        }
                        var v40 int32 = 0
                        _ = v37
                        _ = v40
                    }
                    var v41 uint32 = 2147483648
                    v42 := []int{}
                    v43 := GenS4{F0: v7, F1: -1}
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = v23
            case 4294967295:
                v44 := []int{}
                _ = v44
            case 143:
                switch (^(fi16)) {
                case -2147483648:
                    v45 := GenS5{F0: 0, F1: 264, F2: 4294967295}
                    _ = v45
                case 1:
                    if v7 {
                        v46 := v2[("" + ("\n\t" + (("\n\t" + "") + ("" + ("foo bar" + ("a" + "hello"))))))]
                        v3 = v3
                        _ = v46
                    }
                case 0:
                    var v47 string = "foo bar"
                    _ = v47
                default:
                    var v48 string = (("a" + "hello") + "\n\t")
                    _ = v48
                }
            }
            v49 := []int{(v8 >> ((18446744073709551615) & 63)), (v1 ^ 981), v8, (-(fi22))}
            _ = fi22
            _ = v49
        }
        _ = fi16
    }
    for fi50 := 0; fi50 < 7; fi50++ {
        fmt.Println(v5, v6)
        v51 := v2["hello"]
        switch (-(v1)) {
        case 32767:
            fmt.Println(fi50, v7)
        }
        _ = fi50
        _ = v51
    }
    var v52 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v7
    _ = v8
}
