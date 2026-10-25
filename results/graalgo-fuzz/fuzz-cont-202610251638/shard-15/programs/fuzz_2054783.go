package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }
func genFn0() bool {
    v6 := []int{127, 32768, 541, 145}
    var v7 string = ("foo bar" + "")
    if false {
        fmt.Println(v7)
        var v8 string = (v7 + (v7 + (v7 + v7)))
        _ = v8
    } else {
        v7 = (("hello" + (v7 + ((v7 + ((v7 + v7) + "")) + "a"))) + ("a" + ("\n\t" + ("a" + ""))))
        v7 = "a"
    }
    var v9 uint8 = 2
    _ = v6
    _ = v7
    _ = v9
    return (true && false)
}
func genFn1(p0 int8, p1 int8) int64 {
    p1 -= p0
    var v10 int64 = 0
    v11 := []int{}
    if true {
        var v12 int64 = v10
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return v10
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int16 = 32767
    var v3 int8 = 42
    var v4 uint32 = 2
    var v5 int16 = (v2 / (((v2 | (v2 % (((-(v2))) | 1)))) | 1))
    v3 = (v3 >> ((18446744073709551615) & 63))
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            fmt.Println(v3, v5)
            fmt.Println(v3, v3, fi14)
            _ = fi14
        }
        v15 := []int{}
        _ = fi13
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v4 = (v4 << ((0) & 63))
        if (3.14159 >= 0.0) {
            v4 = v4
            v5 = v5
        } else {
            v17 := []int{fi16}
            var v18 int8 = (v3 / (((v3 - int8(v5))) | 1))
            var v19 uint64 = 361
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    switch ("" + ((("a" + "") + (("a" + "foo bar") + (("a" + "hello") + "a"))) + "a")) {
    case "a":
        v20 := []int{v1, v1, v1, (v1 - 128)}
        _ = v20
    case "foo bar":
        if false {
            if (191 <= 255) {
                for fi21 := 0; fi21 < 6; fi21++ {
                    fmt.Println(v2)
                    if (!(!false)) {
                        fmt.Println(v5, v1, v1)
                        var v22 int = -2147483648
                        _ = v22
                    }
                    if true {
                        if true {
                            v23 := []int{v1}
                            var v24 bool = (false && (false && (42 != 65535)))
                            _ = v23
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 19; fi25++ {
                            v26 := []int{(fi21 % ((v1) | 1)), 2147483647}
                            _ = fi25
                            _ = v26
                        }
                    }
                    _ = fi21
                }
            } else {
                v4 = (v4 ^ (v4 % ((2147483648) | 1)))
                var v27 int64 = 1
                switch ("foo bar" + "hello") {
                case "foo bar":
                    fmt.Println(v4, v1)
                case "a":
                    var v28 uint8 = 0
                    _ = v28
                default:
                    var v29 string = ("foo bar" + ("\n\t" + ""))
                    _ = v29
                }
                _ = v27
            }
            v30 := []int{128, (^(v1)), v1}
            _ = v30
        }
    case "":
        if false {
            if (42 == 42) {
                v31 := GenS1{F0: 127}
                _ = v31
            } else {
                fmt.Println(v5)
            }
            for fi32 := 0; fi32 < 5; fi32++ {
                v33 := []int{fi32, 0, fi32}
                _ = fi32
                _ = v33
            }
            var v34 uint8 = 2
            _ = v34
        }
    default:
        v35 := []int{(-(v1)), 9223372036854775807, v1, 687}
        _ = v35
    }
    v2 -= v2
    v36 := GenS1{F0: uint8(v4)}
    var v37 uint16 = 904
    v1 = (-(v1))
    for fi38 := 0; fi38 < 15; fi38++ {
        for fi39 := 0; fi39 < 3; fi39++ {
            v37 = (v37 | (v37 ^ v37))
            _ = fi39
        }
        for fi40 := 0; fi40 < 1; fi40++ {
            fmt.Println(fi40)
            _ = fi40
        }
        fmt.Println(v5, v5)
        _ = fi38
    }
    v37 -= v37
    v41 := []int{(v1 & (-(v1)))}
    v42 := GenS1{F0: 11}
    for fi43 := 0; fi43 < 18; fi43++ {
        switch ("\n\t" + "hello") {
        case "hello":
            var v44 uint16 = v37
            _ = v44
        case "a":
            v2 += -32768
        case "foo bar":
            if false {
                for fi45 := 0; fi45 < 5; fi45++ {
                    var v46 uint64 = 4294967295
                    _ = fi45
                    _ = v46
                }
                for fi47 := 0; fi47 < 5; fi47++ {
                    if false {
                        fmt.Println(v1)
                        v4 += 976
                    }
                    _ = fi47
                }
            } else {
                v41 = append(v41, (fi43 << ((180) & 63)))
            }
        }
        v2 = v5
        if ((v4 % ((v4) | 1)) == 98) {
            v41 = append(v41, (fi43 >> ((0) & 63)))
        }
        _ = fi43
    }
    var v48 int = (v1 | 42)
    if false {
        var v49 bool = true
        _ = v49
    }
    v5 = 0
    switch (v1 << ((113) & 63)) {
    case -1:
        v37 = v37
    case -590:
        if false {
            v50 := GenS2{F0: (^(v4)), F1: 255, F2: (^(v37))}
            v5 -= (v5 % ((v2) | 1))
            if len(v41) > 0 {
                v41[0] = (v48 ^ (^(v1)))
            }
            _ = v50
        } else {
            var v51 string = "foo bar"
            var v52 int16 = v2
            _ = v51
            _ = v52
        }
    case 930:
        v4 = 2147483648
    default:
        if len(v41) >= 2 {
            v41 = v41[1:len(v41)]
        }
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v48
    _ = v5
}
