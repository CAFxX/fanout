package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 string; F2 uint }
type GenS4 struct { F0 int32 }
func genFn0() uint {
    v7 := GenS1{F0: 231, F1: ""}
    v8 := GenS2{F0: "a"}
    var v9 float64 = 1.5
    v10 := []int{-9223372036854775808, 127, 127, 738}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 628
}
func genFn1() int {
    var v11 int16 = -802
    v12 := GenS2{F0: (("foo bar" + "foo bar") + (((("foo bar" + ("a" + "foo bar")) + "") + ((("" + ("hello" + ((("a" + (((("\n\t" + "a") + (((("foo bar" + "\n\t") + ("\n\t" + ("" + "hello"))) + "a") + "")) + "hello") + "")) + "\n\t") + ""))) + (("" + "a") + "\n\t")) + ("hello" + (("\n\t" + "a") + "\n\t")))) + "hello"))}
    fmt.Println(v11)
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]int)
    var v3 uint8 = 75
    v4 := make(map[uint32]uint32)
    v5 := []string{"a"}
    var v6 int64 = 704
    var v13 int64 = -1
    v6 = v6
    if false {
        var v14 int16 = 42
        if (104 <= 127) {
            if (-46 > 42) {
                switch 198 {
                case -113:
                    v13 += 0
                case 42:
                    for fi15 := 0; fi15 < 12; fi15++ {
                        if ((-106 >= 107) && true) {
                            var v16 float32 = 1e10
                            var v17 uint = 65535
                            v18 := GenS1{F0: (v3 + v3), F1: ("a" + ("hello" + ""))}
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            fmt.Println(v6, v1)
                            var v19 int64 = 1
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 3; fi20++ {
                            v3 += (v3 % ((0) | 1))
                            var v21 int16 = 42
                            delete(v2, "\n\t")
                            _ = fi20
                            _ = v21
                        }
                        _ = fi15
                    }
                }
                switch (("hello" + "") + (("foo bar" + (((("\n\t" + "hello") + "hello") + "") + "")) + "\n\t")) {
                case "foo bar":
                    fmt.Println(v6, v6)
                case "hello":
                    var v22 int = -359
                    _ = v22
                }
            } else {
                var v23 string = ""
                _ = v23
            }
        }
        fmt.Println(v14)
        _ = v14
    } else {
        v24 := GenS1{F0: 127, F1: (((("foo bar" + "") + "") + "foo bar") + "foo bar")}
        fmt.Println(v13, v13)
        fmt.Println(v3, v1, v1)
        _ = v24
    }
    var v25 int8 = 42
    var v26 int = -9223372036854775808
    fmt.Println(v1, v1)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v27 := GenS3{F0: v25, F1: "", F2: genFn0()}
    if (true || true) {
        var v28 uint16 = 42
        v25 = (-(v25))
        _ = v28
    } else {
        if ((("" + ("foo bar" + "")) < "\n\t") || (2 == uint64(v3))) {
            v5 = append(v5, "")
            v29 := v4[703]
            _ = v29
        }
        var v30 int32 = 42
        v30 = v30
        _ = v30
    }
    fmt.Println(v1, v26)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v31 := GenS1{F0: v3, F1: "hello"}
    v32 := GenS2{F0: (("a" + ((("a" + "a") + "hello") + "")) + "")}
    var v33 int64 = (v13 >> ((18446744073709551615) & 63))
    var v34 float64 = 0.0
    v35 := GenS4{F0: -2147483648}
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
