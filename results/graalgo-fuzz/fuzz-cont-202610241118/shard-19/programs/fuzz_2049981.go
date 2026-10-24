package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int64; F1 bool; F2 uint64 }
func genFn0(p0 string) uint64 {
    p0 = p0
    v7 := []int{0}
    var v8 int64 = 723
    if true {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        if len(v7) > 0 {
            v7[0] = 827
        }
        v9 := GenS1{F0: 128}
        _ = v9
    }
    var v10 uint64 = 0
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    return v10
}
func genFn1() string {
    if ("a" < "") {
        var v11 int64 = -956
        fmt.Println(v11)
        _ = v11
    } else {
        v12 := GenS1{F0: 2}
        for fi13 := 0; fi13 < 7; fi13++ {
            if false {
                switch (fi13 - fi13) {
                case 4294967295:
                    fmt.Println(fi13)
                case -704:
                    for fi14 := 0; fi14 < 1; fi14++ {
                        var v15 int = fi13
                        v15 += 872
                        var v16 int64 = 445
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                default:
                    fmt.Println(fi13)
                }
            }
            if true {
                var v17 int16 = -32768
                _ = v17
            }
            _ = fi13
        }
        _ = v12
    }
    var v18 float64 = -2.25
    v18 -= (3.14159 - ((0.1 + float64(0.1)) - (v18 / 1.0)))
    switch ("\n\t" + "\n\t") {
    case "":
        v19 := []int{32768, 9223372036854775807, -483}
        _ = v19
    default:
        if false {
            var v20 uint32 = 548
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 int64 = 545
                if true {
                    var v23 int16 = 0
                    v24 := GenS2{F0: -2147483648, F1: 42}
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(v20, v20)
                    v20 += v20
                    fmt.Println(v18)
                }
                v25 := GenS3{F0: v22, F1: false, F2: 1}
                _ = fi21
                _ = v22
                _ = v25
            }
            fmt.Println(v20, v18)
            _ = v20
        }
    }
    v26 := []int{-209}
    _ = v18
    _ = v26
    return ("" + "")
}

func main() {
    var v1 uint16 = 928
    var v2 int = -1
    v3 := []int{9223372036854775807, int(v1)}
    var v4 float64 = 0.1
    v5 := []string{"a", ("\n\t" + ("hello" + "hello"))}
    var v6 float32 = -2.25
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = (v6 / float32((1.5 + v4)))
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 uint16 = 65535
        v6 += v6
        if (!(true && (191 < 0))) {
            var v29 bool = false
            v30 := GenS2{F0: 0, F1: 154}
            fmt.Println(v29, v29)
            _ = v29
            _ = v30
        } else {
            v31 := GenS3{F0: 42, F1: false, F2: 4294967295}
            switch (fi27 * (v2 & (fi27 / ((4294967295) | 1)))) {
            case -9223372036854775808:
                v32 := GenS1{F0: 102}
                _ = v32
            case 2147483647:
                v4 = 1.5
            default:
                for fi33 := 0; fi33 < 10; fi33++ {
                    v2 += fi27
                    _ = fi33
                }
            }
            _ = v31
        }
        _ = fi27
        _ = v28
    }
    var v34 int64 = -9223372036854775808
    var v35 float32 = 1e10
    var v36 uint = 42
    if len(v5) > 0 {
        v5[0] = (("foo bar" + "hello") + "a")
    }
    var v37 uint8 = 167
    if len(v3) > 0 {
        v3[0] = (v2 + v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
