package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 int16; F2 int64 }
type GenS5 struct { F0 string; F1 int32; F2 string }
type GenS6 struct { F0 int64; F1 int32; F2 uint32 }
type GenS7 struct { F0 string; F1 uint64; F2 int64 }
type GenS8 struct { F0 int8; F1 uint16; F2 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16; F1 int }
func genFn0(p0 uint64, p1 int16) int64 {
    var v7 bool = (true || (!false))
    if v7 {
        p1 -= 32767
        for fi8 := 0; fi8 < 10; fi8++ {
            switch -733 {
            case -671:
                v9 := GenS1{F0: p1, F1: 212}
                _ = v9
            case 127:
                fmt.Println(p0)
            }
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := GenS2{F0: int16(fi8)}
                v12 := GenS2{F0: (p1 | p1)}
                var v13 uint8 = 255
                if v7 {
                    fmt.Println(v13, v13)
                    if (0 != 703) {
                        var v14 int32 = 42
                        v7 = ((float32(1e10) >= 3.14159) && (0.0 <= ((3.14159 * -2.25) - -2.25)))
                        var v15 string = "a"
                        _ = v14
                        _ = v15
                    }
                } else {
                    v16 := GenS3{F0: 4294967295}
                    v7 = v7
                    _ = v16
                }
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            v17 := []int{}
            _ = fi8
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        v7 = (((false && (!v7)) && ((42 >= 0) && true)) && (-2.25 > 3.14159))
        v19 := []int{128, (fi18 - (^(fi18))), (^(fi18))}
        v7 = (("\n\t" + "\n\t") > ("a" + (("" + "") + ((("" + "hello") + "foo bar") + "foo bar"))))
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 int) uint {
    var v20 uint32 = 42
    switch p0 {
    case 127:
        v21 := []int{p0, p0, 32768, p0}
        _ = v21
    case -9223372036854775808:
        switch p0 {
        case -495:
            var v22 float32 = float32((0.0 - float64(-2.25)))
            _ = v22
        case 42:
            if true {
                switch (p0 ^ -9223372036854775808) {
                case 32767:
                    v23 := GenS4{F0: -9223372036854775808, F1: 429, F2: genFn0(4294967295, -32768)}
                    _ = v23
                default:
                    v20 -= (v20 % ((v20) | 1))
                }
                v24 := GenS5{F0: "hello", F1: 2147483647, F2: "a"}
                _ = v24
            } else {
                switch ("foo bar" + "hello") {
                case "hello":
                    fmt.Println(p0)
                case "a":
                    var v25 int = p0
                    _ = v25
                }
            }
        case 9223372036854775807:
            if true {
                for fi26 := 0; fi26 < 14; fi26++ {
                    for fi27 := 0; fi27 < 19; fi27++ {
                        v28 := GenS6{F0: 0, F1: -1, F2: 2}
                        _ = fi27
                        _ = v28
                    }
                    fmt.Println(fi26, v20, p0)
                    _ = fi26
                }
                if (false || false) {
                    v29 := GenS2{F0: 1}
                    p0 = p0
                    var v30 int = -2147483648
                    _ = v29
                    _ = v30
                } else {
                    var v31 int32 = -2147483648
                    _ = v31
                }
            } else {
                var v32 uint8 = 127
                _ = v32
            }
        }
    case 32768:
        for fi33 := 0; fi33 < 10; fi33++ {
            v34 := GenS7{F0: ("a" + "hello"), F1: 959, F2: 42}
            switch ("hello" + "hello") {
            case "":
                fmt.Println(fi33, v20)
            case "a":
                v35 := GenS8{F0: 42, F1: 949, F2: "hello"}
                _ = v35
            default:
                var v36 int8 = -1
                _ = v36
            }
            var v37 uint16 = 255
            _ = fi33
            _ = v34
            _ = v37
        }
    default:
        v20 = v20
    }
    _ = p0
    _ = v20
    return 902
}

func main() {
    var v1 uint8 = 16
    var v2 int64 = 827
    v3 := []string{"a"}
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    var v6 int8 = -1
    if false {
        v2 += (v2 - v2)
    } else {
        v38 := GenS9{F0: false}
        _ = v38
    }
    if ((-2.25 + 3.14159) < 3.14159) {
        for fi39 := 0; fi39 < 6; fi39++ {
            var v40 int = (^(fi39))
            switch ((("a" + "a") + "\n\t") + "foo bar") {
            case "a":
                v1 -= (v1 - (^(v1)))
            case "foo bar":
                fmt.Println(v40, v40)
            case "hello":
                fmt.Println(v6, fi39)
            }
            if false {
                if true {
                    v41 := v5["foo bar"]
                    _ = v41
                } else {
                    if true {
                        fmt.Println(v40, fi39)
                        fmt.Println(v1, v6)
                        v40 += fi39
                    }
                    v42 := GenS10{F0: -1, F1: v40}
                    _ = v42
                }
                fmt.Println(v2)
            }
            _ = fi39
            _ = v40
        }
        v2 = (v2 % ((1) | 1))
        fmt.Println(v1, v1, v1)
    }
    v1 += (^(v1))
    fmt.Println(v1, v6, v2)
    v43 := GenS2{F0: -615}
    v44 := GenS8{F0: v6, F1: 65535, F2: ((((("a" + (("hello" + ("hello" + "a")) + ((("foo bar" + (("foo bar" + ("a" + (("a" + "\n\t") + ((("" + "") + "hello") + "a")))) + "a")) + "hello") + ("\n\t" + "foo bar")))) + "") + ("foo bar" + ("\n\t" + "a"))) + "hello") + "\n\t")}
    var v45 bool = (32767 > -208)
    var v46 uint16 = 1
    for fi47 := 0; fi47 < 14; fi47++ {
        v46 = (v46 / (((v46 % (((^(v46))) | 1))) | 1))
        fmt.Println(fi47, v1)
        var v48 uint = 18446744073709551615
        _ = fi47
        _ = v48
    }
    var v49 uint8 = (v1 ^ uint8(v46))
    v50 := GenS9{F0: false}
    var v51 int8 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
