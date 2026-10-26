package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int64) uint {
    for fi6 := 0; fi6 < 4; fi6++ {
        p0 -= p0
        if true {
            var v7 uint8 = 128
            v8 := []int{fi6}
            for fi9 := 0; fi9 < 20; fi9++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                if true {
                    v10 := GenS1{F0: -128}
                    v11 := GenS1{F0: 127}
                    var v12 int = (fi6 & fi6)
                    _ = v10
                    _ = v11
                    _ = v12
                }
                fmt.Println(p0, fi6)
                var v13 float32 = (3.14159 / 1e10)
                _ = fi9
                _ = v13
            }
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    v14 := []int{}
    var v15 string = "hello"
    v16 := GenS1{F0: 49}
    if true {
        v15 = v15
    } else {
        v17 := GenS1{F0: -1}
        var v18 int64 = 9223372036854775807
        var v19 string = (v15 + v15)
        _ = v17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return 65535
}

func main() {
    var v1 uint32 = 1
    var v2 bool = false
    var v3 int = 42
    var v4 uint = 4294967295
    var v5 int = (^(v3))
    switch ("" + "") {
    case "foo bar":
        if true {
            switch (("" + "a") + "foo bar") {
            case "hello":
                var v20 bool = (!v2)
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 15; fi21++ {
                    if (((2147483647 > 308) && v2) || ((v1 / ((4294967295) | 1)) > v1)) {
                        v22 := []int{128, 42}
                        for fi23 := 0; fi23 < 16; fi23++ {
                            v5 = fi21
                            if v2 {
                                var v24 int = (v3 ^ (^(v3)))
                                _ = v24
                            } else {
                                v25 := GenS1{F0: 42}
                                v3 = (fi21 % ((fi21) | 1))
                                v2 = (!(true || (0 >= 180)))
                                v3 -= int(v4)
                                _ = len(v22)
                                _ = v25
                            }
                            _ = fi23
                        }
                        var v26 uint32 = (v1 | 42)
                        _ = v22
                        _ = v26
                    } else {
                        var v27 int32 = 2147483647
                        _ = v27
                    }
                    var v28 bool = false
                    v3 += (v5 << (((^(v4))) & 63))
                    _ = fi21
                    _ = v28
                }
            case "a":
                var v29 uint8 = 2
                _ = v29
            }
            var v30 uint32 = v1
            _ = v30
        }
    case "\n\t":
        v31 := []int{(v3 >> ((v4) & 63)), int(v4), 9223372036854775807, int(v1)}
        _ = v31
    default:
        v32 := GenS2{F0: 243, F1: 413, F2: v2}
        _ = v32
    }
    if false {
        var v33 uint64 = 2
        v34 := GenS1{F0: 0}
        _ = v33
        _ = v34
    }
    v35 := []int{42, v3}
    v5 = (^(v5))
    v5 = v5
    if len(v35) > 0 {
        v35[0] = v5
    }
    _ = len(v35)
    if (!true) {
        var v36 uint64 = 4294967295
        switch ((("foo bar" + "foo bar") + "\n\t") + ("" + ("hello" + (("" + "hello") + "hello")))) {
        case "foo bar":
            v4 += (v4 * (v4 + uint(v36)))
        case "a":
            if len(v35) > 0 {
                v35[0] = (v3 & -2147483648)
            }
        }
        switch "hello" {
        case "a":
            v37 := GenS3{F0: -1}
            _ = v37
        default:
            var v38 int8 = -128
            _ = v38
        }
        _ = v36
    } else {
        if false {
            v39 := GenS4{F0: ("foo bar" + ""), F1: (false && v2)}
            _ = v39
        }
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
