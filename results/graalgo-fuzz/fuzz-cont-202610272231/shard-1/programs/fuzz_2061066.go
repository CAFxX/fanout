package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int; F1 string; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint }
func genFn0() int32 {
    switch 9223372036854775807 {
    case 32768:
        switch 32767 {
        case 9223372036854775807:
            v4 := []int{138, -1, -9223372036854775808, 2147483647}
            _ = v4
        case 855:
            v5 := GenS1{F0: 255, F1: 2147483647}
            _ = v5
        default:
            var v6 uint64 = 2
            _ = v6
        }
    case 9223372036854775807:
        v7 := GenS2{F0: 127, F1: ("hello" + ("a" + "foo bar")), F2: 42}
        _ = v7
    case -2147483648:
        v8 := []int{}
        _ = v8
    }
    var v9 uint8 = 128
    v10 := GenS1{F0: 706, F1: -2147483648}
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS3{F0: 42}
        v9 -= v9
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v9
    return 1
}
func genFn1() uint {
    var v13 bool = (-2147483648 <= -171)
    v13 = false
    v13 = v13
    _ = v13
    return 18446744073709551615
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []uint32{0}
    var v3 uint8 = 178
    if false {
        var v14 int32 = -76
        _ = v14
    } else {
        var v15 int32 = genFn0()
        v15 -= (v15 + 42)
        v16 := GenS1{F0: 245, F1: v15}
        _ = v15
        _ = v16
    }
    var v17 uint8 = (v3 - 1)
    var v18 uint32 = 4294967295
    var v19 uint32 = v18
    var v20 uint64 = 18446744073709551615
    v19 -= v18
    v17 += (v3 << ((255) & 63))
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 string = ""
            var v24 int16 = 0
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    fmt.Println(v20, v17, v18)
    fmt.Println(v20, v18, v1)
    switch (("foo bar" + "\n\t") + ("hello" + ("\n\t" + "\n\t"))) {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        for fi25 := 0; fi25 < 10; fi25++ {
            if (genFn1() <= 18446744073709551615) {
                if false {
                    v26 := GenS4{F0: v20}
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 int8 = 126
                        var v29 uint = genFn1()
                        if true {
                            var v30 int32 = genFn0()
                            _ = v30
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    _ = v26
                } else {
                    if len(v2) > 0 {
                        v2[0] = v19
                    }
                }
            } else {
                v3 += (^(v3))
                if false {
                    v1 -= (v1 - v1)
                    if (!(!(("\n\t" + (("foo bar" + ("hello" + "a")) + "")) != ""))) {
                        _ = len(v2)
                        var v31 uint8 = 42
                        fmt.Println(v20, v19, fi25)
                        _ = v31
                    } else {
                        var v32 int64 = -9223372036854775808
                        _ = v32
                    }
                    switch (("" + "foo bar") + "") {
                    case "hello":
                        fmt.Println(v17)
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = (v18 >> ((877) & 63))
                        }
                    case "foo bar":
                        var v33 uint32 = 0
                        _ = v33
                    default:
                        var v34 uint8 = 127
                        _ = v34
                    }
                }
            }
            _ = fi25
        }
    case "foo bar":
        v2 = append(v2, 0)
    default:
        for fi35 := 0; fi35 < 13; fi35++ {
            for fi36 := 0; fi36 < 10; fi36++ {
                v1 -= (fi36 % (((fi35 + (fi36 >> ((4294967295) & 63)))) | 1))
                if (v18 == (v18 & 2147483647)) {
                    v37 := GenS4{F0: v20}
                    v38 := GenS1{F0: 1, F1: -2147483648}
                    _ = v37
                    _ = v38
                }
                v2 = append(v2, v19)
                var v39 int8 = 127
                _ = fi36
                _ = v39
            }
            _ = fi35
        }
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        v41 := GenS5{F0: int32(v3), F1: 2}
        var v42 float64 = 3.14159
        v2 = append(v2, (v18 << ((2) & 63)))
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println(v17, v20)
    for fi43 := 0; fi43 < 2; fi43++ {
        if len(v2) > 0 {
            v2[0] = v19
        }
        _ = fi43
    }
    var v44 int32 = 1
    v17 -= (^(v3))
    v45 := GenS5{F0: v44, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v44
    _ = v45
}
