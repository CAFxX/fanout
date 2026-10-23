package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }
type GenS3 struct { F0 int32; F1 int8; F2 uint64 }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0(p0 bool, p1 uint64, p2 uint32) uint32 {
    p2 -= p2
    if (!p0) {
        p0 = p0
    }
    if p0 {
        v4 := GenS1{F0: 42, F1: -90, F2: 42}
        _ = v4
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 string = ("a" + ("foo bar" + "a"))
    v2 := make(map[int]bool)
    var v3 int16 = 32767
    for fi5 := 0; fi5 < 5; fi5++ {
        switch v1 {
        case "":
            v6 := GenS1{F0: -1, F1: -1, F2: 255}
            _ = v6
        case "foo bar":
            for fi7 := 0; fi7 < 4; fi7++ {
                fmt.Println(v1, v1, v1)
                _ = len(v2)
                v8 := GenS2{F0: 1, F1: 1, F2: ((v3 - v3) == (-(v3)))}
                if true {
                    var v9 int32 = 46
                    switch ("a" + "hello") {
                    case "hello":
                        v9 = v9
                    case "a":
                        var v10 uint = 486
                        _ = v10
                    case "foo bar":
                        v11 := GenS2{F0: 0, F1: 42, F2: (32767 < 2)}
                        _ = v11
                    }
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
        default:
            for fi12 := 0; fi12 < 12; fi12++ {
                v3 = (v3 / ((-32768) | 1))
                _ = fi12
            }
        }
        fmt.Println(v3)
        fmt.Println(v3, v1)
        _ = fi5
    }
    v13 := GenS3{F0: 0, F1: 42, F2: 4294967295}
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int = (fi14 / ((1) | 1))
        _ = fi14
        _ = v15
    }
    v1 = v1
    var v16 string = v1
    switch -884 {
    case 1:
        for fi17 := 0; fi17 < 13; fi17++ {
            if (v1 <= v1) {
                var v18 uint16 = 1
                var v19 uint = 42
                _ = v18
                _ = v19
            }
            v20 := v2[fi17]
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 uint16 = 65535
                if (!(("hello" + ("\n\t" + v1)) < "hello")) {
                    var v23 uint8 = 53
                    v1 = v16
                    delete(v2, (fi21 + (fi17 >> ((255) & 63))))
                    _ = v23
                } else {
                    var v24 int64 = 740
                    if true {
                        v20 = true
                        fmt.Println(v3, v24)
                    } else {
                        switch int(v3) {
                        case 2147483647:
                            var v25 string = "a"
                            _ = v25
                        case -94:
                            var v26 bool = v20
                            _ = v26
                        }
                    }
                    _ = v24
                }
                var v27 float32 = 0.0
                _ = fi21
                _ = v22
                _ = v27
            }
            _ = fi17
            _ = v20
        }
    case -9223372036854775808:
        v28 := GenS4{F0: v3, F1: -1}
        _ = v28
    case 424:
        v29 := []int{-2147483648}
        _ = v29
    }
    v16 = v1
    var v30 int64 = 0
    if (!false) {
        fmt.Println(v3, v1)
        var v31 int8 = 127
        _ = v31
    } else {
        if true {
            v3 -= 1
        }
        if false {
            switch (v16 + (((((v1 + (v1 + "foo bar")) + "hello") + (v1 + ("" + ("hello" + (v16 + v16))))) + v1) + v1)) {
            case "":
                if true {
                    switch 127 {
                    case 4294967295:
                        v1 = ((v16 + v16) + v1)
                    }
                    var v32 uint64 = 169
                    _ = v32
                }
            default:
                var v33 uint64 = 1
                _ = v33
            }
            var v34 float32 = 3.14159
            v2[-1] = (0 == 42)
            _ = v34
        } else {
            switch v1 {
            case "\n\t":
                v35 := []int{86, 127}
                _ = v35
            case "a":
                if (627 != 18446744073709551615) {
                    switch v1 {
                    case "":
                        for fi36 := 0; fi36 < 9; fi36++ {
                            v37 := []int{}
                            _ = len(v2)
                            _ = fi36
                            _ = v37
                        }
                    default:
                        fmt.Println(v1)
                    }
                }
            default:
                v30 = v30
            }
            var v38 uint32 = 2147483648
            _ = v38
        }
        fmt.Println(v3, v1)
    }
    v16 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v30
}
