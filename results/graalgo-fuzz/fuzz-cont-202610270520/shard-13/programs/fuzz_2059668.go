package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
func genFn0() uint16 {
    if (1 == -9223372036854775808) {
        for fi4 := 0; fi4 < 6; fi4++ {
            fmt.Println(fi4)
            v5 := GenS1{F0: true}
            var v6 uint32 = 1
            _ = fi4
            _ = v5
            _ = v6
        }
        var v7 int = -9223372036854775808
        _ = v7
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            switch 32768 {
            case 128:
                fmt.Println(fi8)
            case 32768:
                if (!false) {
                    switch (fi8 | 32768) {
                    case 32768:
                        var v9 uint = 255
                        _ = v9
                    }
                    var v10 int32 = -1
                    _ = v10
                }
            default:
                var v11 float32 = 0.1
                _ = v11
            }
            v12 := GenS1{F0: ((-1 < -133) || false)}
            _ = fi8
            _ = v12
        }
    }
    var v13 string = ("foo bar" + "hello")
    _ = v13
    return 65535
}

func main() {
    var v1 int64 = 1
    var v2 string = "\n\t"
    var v3 bool = (-2147483648 > -1)
    v2 = v2
    if (v2 > "\n\t") {
        var v14 string = "a"
        var v15 float32 = ((-2.25 * (0.1 + (((((-2.25 / ((1.5 / 1e10) * (1.5 - 0.0))) + 1.5) / 0.1) - 1e10) + (0.0 * 1e10)))) + 1e10)
        switch (("" + v14) + v14) {
        case "hello":
            v14 = "foo bar"
        case "a":
            for fi16 := 0; fi16 < 17; fi16++ {
                v1 = 529
                _ = fi16
            }
        case "\n\t":
            var v17 float32 = (v15 / v15)
            _ = v17
        }
        _ = v14
        _ = v15
    } else {
        for fi18 := 0; fi18 < 3; fi18++ {
            switch fi18 {
            case 0:
                fmt.Println(v3)
            default:
                for fi19 := 0; fi19 < 13; fi19++ {
                    v2 = (v2 + v2)
                    _ = fi19
                }
            }
            v20 := []int{fi18}
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 float32 = (0.0 + 1e10)
                v22 = 0.1
                if (true && (1 <= 734)) {
                    v23 := GenS1{F0: v3}
                    v1 = (v1 % ((v1) | 1))
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
            _ = fi18
            _ = v20
        }
    }
    var v24 float64 = 3.14159
    var v25 uint32 = 1
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 int = (fi26 | -502)
        v28 := []int{(fi26 ^ -588), fi26}
        v3 = v3
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: v3}
    for fi30 := 0; fi30 < 8; fi30++ {
        if v3 {
            fmt.Println(v25)
        }
        if (fi30 < (fi30 & (fi30 % (((^(fi30))) | 1)))) {
            if (107 == 42) {
                switch "a" {
                case "\n\t":
                    var v31 int32 = 1
                    _ = v31
                case "a":
                    for fi32 := 0; fi32 < 8; fi32++ {
                        v1 += v1
                        _ = fi32
                    }
                case "foo bar":
                    if (v2 <= v2) {
                        var v33 int8 = 52
                        if v3 {
                            v25 = 2147483648
                            if v3 {
                                var v34 float64 = (v24 * v24)
                                v35 := GenS1{F0: v3}
                                _ = v34
                                _ = v35
                            }
                        } else {
                            var v36 uint64 = 0
                            _ = v36
                        }
                        _ = v33
                    } else {
                        v37 := GenS2{F0: 4294967295}
                        _ = v37
                    }
                }
                if (v3 && true) {
                    v25 -= (v25 ^ (v25 - (v25 % ((4294967295) | 1))))
                } else {
                    var v38 int64 = -9223372036854775808
                    _ = v38
                }
            }
        } else {
            switch (v2 + v2) {
            case "\n\t":
                var v39 float32 = 0.0
                _ = v39
            default:
                var v40 uint32 = (v25 & 647)
                _ = v40
            }
        }
        _ = fi30
    }
    var v41 int16 = 1000
    var v42 uint64 = 4294967295
    fmt.Println(v25)
    switch -1 {
    case -960:
        var v43 uint = 42
        _ = v43
    case -9223372036854775808:
        fmt.Println(v25, v1)
    case 4294967295:
        v3 = (!((727 >= 42) || (!false)))
    }
    v1 -= 9223372036854775807
    for fi44 := 0; fi44 < 9; fi44++ {
        v41 += 0
        v45 := []int{fi44, fi44, fi44}
        _ = fi44
        _ = v45
    }
    fmt.Println(v3, v1, v3)
    if (4294967295 != 2) {
        var v46 uint = 18446744073709551615
        v47 := []int{}
        var v48 uint16 = genFn0()
        _ = v46
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v41
    _ = v42
}
