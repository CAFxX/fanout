package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int8; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
func genFn0() int {
    var v5 int64 = 42
    if true {
        fmt.Println(v5)
        var v6 int8 = 100
        _ = v6
    } else {
        var v7 int32 = 258
        v7 = (v7 * 0)
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{}
        var v10 bool = true
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: "foo bar"}
    _ = v11
    _ = v5
    return 1
}
func genFn1(p0 uint, p1 uint) string {
    switch (((((("a" + "foo bar") + ((("a" + ("\n\t" + "hello")) + ("hello" + "foo bar")) + "\n\t")) + "") + "") + ("hello" + (((("" + "\n\t") + "a") + ("a" + "hello")) + (("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))) + "")))) + ("foo bar" + "hello")) {
    case "a":
        var v12 uint32 = 4294967295
        _ = v12
    }
    v13 := []int{0, -802}
    p0 = (p0 | 567)
    switch "\n\t" {
    case "":
        if len(v13) > 0 {
            v13[0] = -981
        }
    case "\n\t":
        var v14 uint64 = 4294967295
        _ = v14
    default:
        v15 := GenS1{F0: "hello"}
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 int = (fi16 + 182)
        var v18 uint64 = 844
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v13
    return "\n\t"
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{1}
    v3 := []uint32{}
    var v4 uint16 = 84
    switch ("foo bar" + ("\n\t" + "foo bar")) {
    case "hello":
        v19 := GenS2{F0: 1}
        _ = v19
    default:
        switch -1 {
        case 959:
            switch ("" + "") {
            case "":
                v20 := GenS3{F0: 272, F1: -63, F2: 0}
                _ = v20
            }
        default:
            for fi21 := 0; fi21 < 5; fi21++ {
                fmt.Println(fi21, fi21)
                if (true || (true || false)) {
                    var v22 uint16 = (v4 ^ v4)
                    if true {
                        if (42 >= 1) {
                            if true {
                                v3 = append(v3, 444)
                            }
                            var v23 uint64 = 2
                            var v24 int8 = 127
                            _ = v23
                            _ = v24
                        } else {
                            v25 := GenS2{F0: 59}
                            _ = v25
                        }
                        v26 := GenS3{F0: 255, F1: 127, F2: 77}
                        _ = v26
                    } else {
                        fmt.Println(v22)
                        if (42 >= 2147483648) {
                            delete(v1, ("\n\t" + "a"))
                            var v27 uint = 2
                            var v28 int16 = -743
                            _ = v27
                            _ = v28
                        } else {
                            v4 = v4
                            v29 := v1["hello"]
                            _ = v29
                        }
                        var v30 float32 = 1.5
                        _ = v30
                    }
                    _ = v22
                }
                _ = fi21
            }
        }
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        v4 += v4
        _ = fi31
    }
    v32 := GenS4{F0: false}
    if (0 != -9223372036854775808) {
        var v33 int64 = 1
        for fi34 := 0; fi34 < 1; fi34++ {
            for fi35 := 0; fi35 < 5; fi35++ {
                var v36 uint8 = 214
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
        _ = v33
    } else {
        var v37 int32 = 0
        _ = v37
    }
    var v38 uint64 = 472
    v39 := GenS5{F0: 255}
    v40 := GenS6{F0: 703}
    v41 := GenS1{F0: "\n\t"}
    if true {
        var v42 uint16 = v4
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v42
    } else {
        fmt.Println(v4, v38)
        v3 = append(v3, 26)
    }
    if true {
        v38 -= (v38 / ((v38) | 1))
        for fi43 := 0; fi43 < 19; fi43++ {
            var v44 uint16 = (v4 << ((65535) & 63))
            _ = fi43
            _ = v44
        }
    }
    switch "foo bar" {
    case "":
        v45 := GenS6{F0: 0}
        _ = v45
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
