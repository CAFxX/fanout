package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 bool; F1 string; F2 uint16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint64) string {
    fmt.Println(p0, p1)
    v6 := []int{-292, -148}
    if len(v6) > 0 {
        v6[0] = -1
    }
    _ = p0
    _ = p1
    _ = v6
    return ((("\n\t" + "\n\t") + "a") + (("foo bar" + "") + ("\n\t" + "hello")))
}
func genFn1(p0 int32, p1 int16, p2 uint32) int64 {
    switch "foo bar" {
    case "\n\t":
        if false {
            switch "\n\t" {
            case "\n\t":
                p0 -= p0
            case "hello":
                var v7 bool = true
                _ = v7
            default:
                fmt.Println(p2, p1, p1)
            }
            var v8 uint32 = p2
            _ = v8
        } else {
            p0 = (p0 << ((42) & 63))
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 uint = 860
                _ = fi9
                _ = v10
            }
        }
    default:
        p1 = (p1 >> ((18446744073709551615) & 63))
    }
    var v11 uint = 18446744073709551615
    var v12 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 bool = (!true)
    var v2 uint16 = 947
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    if true {
        if v1 {
            if (!(!true)) {
                var v13 uint16 = (^(v2))
                _ = v13
            } else {
                if v1 {
                    v14 := []int{9223372036854775807}
                    var v15 string = "hello"
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v1 = v1
                        v17 := []int{(fi16 << ((1) & 63)), fi16}
                        _ = fi16
                        _ = v17
                    }
                }
            }
            var v18 uint = 18446744073709551615
            _ = v18
        } else {
            var v19 int64 = 1
            _ = v19
        }
        v20 := GenS1{F0: 42, F1: 0, F2: 997}
        v21 := []int{2147483647, 127}
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v2 = 255
        v23 := GenS2{F0: v1, F1: "foo bar", F2: (v2 + (v2 * (^(v2))))}
        _ = fi22
        _ = v23
    }
    v24 := v5[-330]
    if v24 {
        var v25 uint32 = 2
        _ = v25
    }
    var v26 int8 = 100
    for fi27 := 0; fi27 < 20; fi27++ {
        v24 = (!(fi27 < 1))
        v1 = v1
        _ = fi27
    }
    var v28 uint64 = 0
    var v29 int8 = v26
    if (2 < 255) {
        switch 127 {
        case 42:
            if false {
                fmt.Println(v24, v1, v29)
            } else {
                for fi30 := 0; fi30 < 3; fi30++ {
                    for fi31 := 0; fi31 < 10; fi31++ {
                        fmt.Println(v29)
                        if false {
                            switch "hello" {
                            case "a":
                                v29 = v26
                            }
                            _ = len(v5)
                        }
                        if ((-(fi30)) < fi30) {
                            if (!(0 <= 42)) {
                                var v32 float64 = 1e10
                                var v33 int8 = 0
                                if false {
                                    fmt.Println(v26)
                                    switch ("hello" + "a") {
                                    case "hello":
                                        switch (("hello" + "foo bar") + ((("a" + (("foo bar" + "a") + ("foo bar" + "foo bar"))) + "a") + "foo bar")) {
                                        case "hello":
                                            fmt.Println(v26, fi30)
                                        }
                                    case "":
                                        v34 := GenS1{F0: -1, F1: int16(v28), F2: -1}
                                        _ = v34
                                    case "foo bar":
                                        v29 += int8(v28)
                                    }
                                } else {
                                    fmt.Println(v28)
                                }
                                _ = v32
                                _ = v33
                            }
                            _ = len(v3)
                        }
                        var v35 float32 = 0.1
                        _ = fi31
                        _ = v35
                    }
                    switch (fi30 + (fi30 << ((1) & 63))) {
                    case 2147483647:
                        for fi36 := 0; fi36 < 17; fi36++ {
                            v37 := GenS3{F0: 2}
                            fmt.Println(v29, fi36)
                            v38 := GenS3{F0: v2}
                            v29 -= 0
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    default:
                        var v39 float32 = (-2.25 / (-2.25 - (1e10 - 0.1)))
                        _ = v39
                    }
                    v1 = v1
                    _ = fi30
                }
                v40 := []int{-1, 882, -9223372036854775808, -2147483648}
                fmt.Println(v28, v1)
                _ = v40
            }
        }
    } else {
        for fi41 := 0; fi41 < 2; fi41++ {
            v42 := GenS3{F0: uint16(v28)}
            _ = fi41
            _ = v42
        }
        fmt.Println(v29)
    }
    var v43 float64 = 1.5
    v44 := []int{2147483647}
    v26 = v26
    var v45 float32 = (0.0 + 0.1)
    v44 = append(v44, -1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
