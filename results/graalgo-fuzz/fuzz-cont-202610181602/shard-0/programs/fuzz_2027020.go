package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{42, 4294967295}
    v3 := []string{"", ("\n\t" + (("hello" + ("" + "")) + "hello"))}
    if false {
        for fi4 := 0; fi4 < 15; fi4++ {
            var v5 string = ""
            v6 := GenS1{F0: 2147483647, F1: fi4}
            var v7 uint16 = 761
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    } else {
        var v8 uint16 = 65535
        v8 -= 2
        v8 += (^(v8))
        _ = v8
    }
    switch (("\n\t" + "a") + "a") {
    case "\n\t":
        if true {
            var v9 int64 = 42
            var v10 float64 = 0.1
            _ = v10
            _ = v9
        } else {
            var v11 uint8 = 127
            _ = v11
        }
    }
    var v12 int8 = 127
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 int64 = 1
        if (65535 > 42) {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                var v15 int64 = -1
                _ = v15
            case "hello":
                fmt.Println(v14, v12, v12)
            }
        }
        switch ("a" + (("a" + "foo bar") + (("\n\t" + ("foo bar" + "hello")) + ("hello" + "\n\t")))) {
        case "a":
            var v16 uint64 = 42
            _ = v16
        case "hello":
            delete(v1, 2147483648)
        case "\n\t":
            var v17 uint32 = 196
            _ = v17
        default:
            var v18 int = fi13
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    v12 += -1
    v12 -= (v12 << ((255) & 63))
    v19 := GenS1{F0: 2, F1: 127}
    for fi20 := 0; fi20 < 4; fi20++ {
        if false {
            switch "a" {
            case "hello":
                var v21 float32 = ((0.1 + (0.1 * (0.1 + -2.25))) + -2.25)
                _ = v21
            case "foo bar":
                if false {
                    v22 := GenS1{F0: 2, F1: (fi20 >> ((4294967295) & 63))}
                    _ = v22
                } else {
                    switch "" {
                    case "":
                        switch ("a" + ("" + "hello")) {
                        case "hello":
                            for fi23 := 0; fi23 < 20; fi23++ {
                                fmt.Println(fi20, fi20)
                                var v24 bool = true
                                v24 = (!(18446744073709551615 > 0))
                                _ = fi23
                                _ = v24
                            }
                        case "a":
                            if (2 > 0) {
                                v25 := GenS2{F0: false}
                                for fi26 := 0; fi26 < 7; fi26++ {
                                    if (("\n\t" + "hello") >= "\n\t") {
                                        v3 = append(v3, ("a" + ("foo bar" + "a")))
                                    } else {
                                        var v27 uint = 18446744073709551615
                                        var v28 int16 = -1
                                        var v29 int64 = -1
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                    fmt.Println(v12)
                                    if false {
                                        switch fi20 {
                                        case 32767:
                                            v1[487] = 42
                                        default:
                                            if true {
                                                var v30 uint = 872
                                                _ = v30
                                            } else {
                                                _ = len(v1)
                                            }
                                        }
                                    }
                                    _ = fi26
                                }
                                _ = v25
                            } else {
                                var v31 bool = true
                                v32 := v1[4294967295]
                                fmt.Println(v32, fi20)
                                _ = v31
                                _ = v32
                            }
                        default:
                            v1[189] = 1
                        }
                    case "foo bar":
                        v1[765] = 4294967295
                    case "hello":
                        fmt.Println(v12, fi20)
                    }
                    var v33 int = (fi20 / ((fi20) | 1))
                    _ = v33
                }
            case "":
                var v34 float64 = 0.1
                _ = v34
            }
            var v35 float64 = -2.25
            _ = v35
        } else {
            fmt.Println(fi20, fi20)
            v12 = (v12 | v12)
        }
        _ = fi20
    }
    v36 := GenS2{F0: (255 >= 4294967295)}
    var v37 int = -1
    fmt.Println(v12, v12)
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v38 uint64 = 1
        v12 = v12
        for fi39 := 0; fi39 < 5; fi39++ {
            if ((-2.25 - 1.5) < (1e10 - (1.5 - -2.25))) {
                v12 = (v12 / ((127) | 1))
                for fi40 := 0; fi40 < 15; fi40++ {
                    v1[4294967295] = 65535
                    _ = fi40
                }
                fmt.Println(v38, v38, v38)
            } else {
                var v41 int16 = -956
                _ = v41
            }
            v12 -= v12
            v42 := GenS2{F0: (!false)}
            _ = fi39
            _ = v42
        }
        _ = v38
    }
    var v43 string = ("\n\t" + ("hello" + "\n\t"))
    fmt.Println(v37, v12)
    v44 := GenS1{F0: 775, F1: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v43
    _ = v44
}
