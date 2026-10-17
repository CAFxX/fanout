package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    var v4 int8 = 0
    var v5 int32 = -685
    if true {
        var v6 float32 = 0.0
        fmt.Println(v5, v5, v5)
        _ = v6
    } else {
        if v1 {
            if true {
                v4 += (v4 | v4)
                v7 := GenS1{F0: "\n\t"}
                switch -502 {
                case -2147483648:
                    v4 += (v4 - v4)
                case 32768:
                    v8 := GenS1{F0: ("\n\t" + "hello")}
                    _ = v8
                }
                _ = v7
            }
            if v1 {
                v5 = (v5 / ((v5) | 1))
                v9 := v2[2147483647]
                _ = v9
            }
            v1 = (0 > 76)
        }
        var v10 uint8 = 135
        var v11 uint = 42
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        _ = len(v2)
        _ = fi12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: (((("a" + "\n\t") + "hello") + "hello") + "")}
        _ = fi13
        _ = v14
    }
    var v15 int64 = 73
    for fi16 := 0; fi16 < 14; fi16++ {
        if v1 {
            if v1 {
                for fi17 := 0; fi17 < 15; fi17++ {
                    v15 = v15
                    _ = fi17
                }
                for fi18 := 0; fi18 < 14; fi18++ {
                    switch (fi18 * 42) {
                    case 4294967295:
                        switch "\n\t" {
                        case "a":
                            v1 = (!((2 != 839) || (v5 <= (v5 % ((v5) | 1)))))
                        default:
                            v19 := []int{fi18, (fi16 % (((fi18 | fi18)) | 1))}
                            _ = v19
                        }
                    default:
                        var v20 int64 = 1
                        _ = v20
                    }
                    v5 += 0
                    if v1 {
                        v15 = v15
                        if (4294967295 > 1) {
                            v21 := []int{fi16}
                            _ = v21
                        } else {
                            fmt.Println(v1)
                            fmt.Println(v4, v15)
                            v15 = (v15 + v15)
                        }
                        v22 := GenS1{F0: ("hello" + "foo bar")}
                        if ((false || true) && (("foo bar" + "\n\t") >= (("a" + ("foo bar" + "a")) + "hello"))) {
                            v4 = v4
                            v4 -= (v4 - (v4 ^ v4))
                        } else {
                            v4 = (^(v4))
                            fmt.Println(v1)
                            var v23 uint16 = 1
                            _ = v23
                        }
                        var v24 uint32 = 2147483648
                        var v25 int64 = -208
                        delete(v2, fi16)
                        v26 := v3[fi18]
                        v26 = v26
                        var v27 string = "a"
                        _ = v22
                        _ = v24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = fi18
                }
            }
            var v28 int8 = v4
            for fi29 := 0; fi29 < 12; fi29++ {
                v15 = v15
                v5 -= (v5 / ((v5) | 1))
                _ = fi29
            }
            _ = v28
        }
        _ = fi16
    }
    v30 := []int{32768, -2147483648, -2147483648}
    for fi31 := 0; fi31 < 20; fi31++ {
        if v1 {
            var v32 uint16 = 1
            _ = v32
        } else {
            v4 -= 127
        }
        switch ("a" + "foo bar") {
        case "a":
            var v33 int32 = v5
            _ = v33
        case "\n\t":
            v2[fi31] = ((!((fi31 + fi31) < fi31)) || ((fi31 << ((1) & 63)) >= (-(fi31))))
        }
        v34 := GenS1{F0: "hello"}
        _ = fi31
        _ = v34
    }
    var v35 int = 128
    v1 = v1
    var v36 float32 = 3.14159
    switch "hello" {
    case "foo bar":
        _ = len(v2)
    case "\n\t":
        fmt.Println(v35)
    default:
        var v37 int8 = (v4 * (v4 + (v4 << ((0) & 63))))
        _ = v37
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        fmt.Println(v35, v1)
        _ = fi38
    }
    if v1 {
        v5 -= (v5 | v5)
    }
    if v1 {
        v39 := GenS1{F0: "\n\t"}
        v35 -= -544
        _ = v39
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := GenS1{F0: ""}
            _ = fi41
            _ = v42
        }
        v43 := GenS1{F0: "foo bar"}
        _ = fi40
        _ = v43
    }
    v44 := GenS1{F0: ("hello" + "hello")}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v44
    _ = v5
}
