package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0(p0 uint) int64 {
    v4 := GenS1{F0: 0}
    switch "foo bar" {
    case "\n\t":
        v5 := []int{42, -2147483648}
        _ = v5
    case "":
        v6 := []int{}
        _ = v6
    }
    switch (("" + ("" + ("hello" + (("foo bar" + "\n\t") + ("" + "hello"))))) + ("hello" + (("\n\t" + "") + ""))) {
    case "":
        var v7 bool = true
        _ = v7
    case "hello":
        p0 = 692
    case "a":
        fmt.Println(p0)
    default:
        var v8 uint32 = 839
        _ = v8
    }
    v9 := GenS1{F0: 376}
    _ = p0
    _ = v4
    _ = v9
    return 103
}
func genFn1(p0 uint32, p1 uint8) int16 {
    var v10 float32 = 0.1
    var v11 uint16 = uint16(p0)
    for fi12 := 0; fi12 < 6; fi12++ {
        if true {
            v13 := []int{int(p0), fi12}
            v14 := GenS2{F0: 0, F1: 94}
            p0 = p0
            _ = v13
            _ = v14
        } else {
            v15 := GenS1{F0: -9223372036854775808}
            if ("foo bar" != ("" + "foo bar")) {
                v16 := GenS1{F0: -253}
                if ((p0 | 65535) >= 934) {
                    v10 += v10
                    v10 = 0.1
                    switch (fi12 & fi12) {
                    case 218:
                        for fi17 := 0; fi17 < 9; fi17++ {
                            var v18 uint64 = 4294967295
                            _ = fi17
                            _ = v18
                        }
                    default:
                        switch fi12 {
                        case 42:
                            var v19 string = "hello"
                            _ = v19
                        case -371:
                            p1 = 128
                        case 2147483647:
                            for fi20 := 0; fi20 < 20; fi20++ {
                                p1 += 1
                                v21 := GenS3{F0: 42}
                                if false {
                                    v22 := GenS1{F0: 42}
                                    var v23 int8 = 0
                                    var v24 uint8 = 1
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                if false {
                                    var v25 string = ("foo bar" + "foo bar")
                                    _ = v25
                                }
                                var v26 float32 = 0.1
                                fmt.Println(p1)
                                _ = fi20
                                _ = v21
                                _ = v26
                            }
                        }
                    }
                }
                if true {
                    fmt.Println(fi12, v11, fi12)
                }
                _ = v16
            } else {
                v11 = v11
                v27 := GenS2{F0: 0, F1: 0}
                for fi28 := 0; fi28 < 16; fi28++ {
                    var v29 bool = false
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
            _ = v15
        }
        if (p1 >= p1) {
            v30 := []int{32768, fi12, fi12}
            if true {
                v31 := GenS1{F0: 9223372036854775807}
                _ = v31
            }
            var v32 uint = 5
            _ = v30
            _ = v32
        }
        fmt.Println(p0, fi12, p1)
        _ = fi12
    }
    v33 := []int{127, 128, 951, -2147483648}
    p1 += p1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v33
    return 741
}

func main() {
    var v1 uint8 = 226
    var v2 int = 4294967295
    v3 := make(map[int]string)
    v1 += (^(v1))
    v34 := v3[(v2 % ((v2) | 1))]
    if false {
        for fi35 := 0; fi35 < 7; fi35++ {
            v36 := []int{(fi35 + 4294967295), fi35}
            v1 = v1
            v34 = v34
            _ = fi35
            _ = v36
        }
    }
    v2 = v2
    var v37 int32 = 2147483647
    v38 := []int{(v2 ^ (v2 ^ (v2 / ((v2) | 1)))), (v2 >> ((685) & 63))}
    var v39 int = 1
    if false {
        var v40 uint32 = 42
        v2 += (-(v39))
        _ = v40
    }
    v2 -= v39
    v41 := GenS4{F0: v37, F1: 1}
    v42 := GenS4{F0: (v37 >> ((108) & 63)), F1: 69}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v41
    _ = v42
}
