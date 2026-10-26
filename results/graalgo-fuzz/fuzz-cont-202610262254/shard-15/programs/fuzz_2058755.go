package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint16; F1 int; F2 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 int; F2 string }
func genFn0() bool {
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: "hello"}
        _ = fi6
        _ = v7
    }
    switch "foo bar" {
    case "hello":
        v8 := []int{0, -765, -126, -1}
        _ = v8
    }
    return (2147483648 <= 2147483648)
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = -2.25
    v3 := []string{"a", "a", ("hello" + "hello"), ""}
    var v4 float64 = 3.14159
    v5 := []string{"hello", (((("hello" + "") + "hello") + "") + "\n\t")}
    _ = len(v1)
    v9 := GenS2{F0: 988, F1: -32768}
    var v10 uint8 = 170
    fmt.Println(v4, v10, v10)
    for fi11 := 0; fi11 < 20; fi11++ {
        switch "" {
        case "hello":
            var v12 float64 = v4
            _ = v12
        case "":
            v13 := GenS3{F0: 32767, F1: (fi11 ^ 127), F2: 0}
            _ = v13
        default:
            switch ("\n\t" + "hello") {
            case "":
                switch "foo bar" {
                case "a":
                    var v14 int = (fi11 ^ fi11)
                    _ = v14
                case "\n\t":
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 bool = false
                        _ = fi15
                        _ = v16
                    }
                case "":
                    v4 = (v4 * (v4 + v4))
                }
            case "a":
                var v17 uint16 = 2
                _ = v17
            default:
                if (!("a" != "\n\t")) {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                        var v19 float32 = v2
                        _ = fi18
                        _ = v19
                    }
                    v2 = v2
                    if (!(2 == 4294967295)) {
                        var v20 int64 = 42
                        _ = v20
                    } else {
                        if false {
                            v2 = -2.25
                            if false {
                                delete(v1, fi11)
                                for fi21 := 0; fi21 < 6; fi21++ {
                                    v2 = v2
                                    _ = fi21
                                }
                            } else {
                                var v22 float32 = ((v2 / v2) * float32((v4 + float64(v2))))
                                _ = v22
                            }
                        } else {
                            fmt.Println(fi11, v2)
                        }
                        fmt.Println(v10, v4)
                    }
                }
            }
        }
        if false {
            var v23 int64 = -1
            _ = v23
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi11
    }
    v24 := GenS4{F0: "\n\t"}
    for fi25 := 0; fi25 < 19; fi25++ {
        v2 -= (0.0 / 1.5)
        for fi26 := 0; fi26 < 2; fi26++ {
            v4 -= (1.5 + v4)
            _ = fi26
        }
        _ = fi25
    }
    v5 = append(v5, "a")
    v27 := v1[-9223372036854775808]
    for fi28 := 0; fi28 < 10; fi28++ {
        if false {
            var v29 bool = v27
            _ = v29
        }
        _ = fi28
    }
    var v30 int32 = 0
    v31 := GenS4{F0: "\n\t"}
    v2 += v2
    if v27 {
        switch ("" + "a") {
        case "a":
            for fi32 := 0; fi32 < 14; fi32++ {
                var v33 int = (-(fi32))
                v1[v33] = (4294967295 == 0)
                v34 := GenS3{F0: 580, F1: 127, F2: v30}
                _ = fi32
                _ = v33
                _ = v34
            }
        case "":
            if (v30 != v30) {
                fmt.Println(v30, v2, v27)
            } else {
                switch "\n\t" {
                case "":
                    for fi35 := 0; fi35 < 15; fi35++ {
                        var v36 uint8 = v10
                        fmt.Println(v10, v27)
                        _ = fi35
                        _ = v36
                    }
                }
                for fi37 := 0; fi37 < 8; fi37++ {
                    if (!(v30 > v30)) {
                        v38 := GenS1{F0: "\n\t"}
                        _ = v38
                    } else {
                        for fi39 := 0; fi39 < 5; fi39++ {
                            var v40 uint = 855
                            _ = fi39
                            _ = v40
                        }
                        v41 := GenS4{F0: ""}
                        _ = v41
                    }
                    _ = fi37
                }
            }
        }
        var v42 uint16 = 0
        _ = v42
    } else {
        v3 = append(v3, (("\n\t" + "hello") + ""))
    }
    switch "\n\t" {
    case "\n\t":
        v43 := GenS5{F0: 0, F1: -1, F2: (((("a" + ("" + "foo bar")) + "\n\t") + (("hello" + ("\n\t" + "")) + ("a" + "\n\t"))) + ("hello" + (("foo bar" + "\n\t") + "foo bar")))}
        _ = v43
    default:
        var v44 int8 = 0
        _ = v44
    }
    var v45 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v45
    _ = v5
    _ = v9
}
