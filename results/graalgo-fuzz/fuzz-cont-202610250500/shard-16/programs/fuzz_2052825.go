package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint16, p1 uint16) int16 {
    v5 := []int{1, -2147483648, 403}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "" {
    case "":
        if false {
            fmt.Println(p1)
            switch ("" + ("\n\t" + "\n\t")) {
            case "a":
                var v6 uint = 18446744073709551615
                _ = v6
            case "foo bar":
                p1 -= p0
            }
        }
    case "\n\t":
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 12; fi8++ {
                p1 -= 179
                v9 := GenS1{F0: (p1 ^ (p1 + p1)), F1: 1, F2: ("foo bar" + "a")}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case "foo bar":
        v10 := GenS1{F0: p1, F1: 18446744073709551615, F2: "hello"}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    return 32767
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 2147483647
    var v3 bool = (!(!(((false || (0 >= 2147483647)) && true) && (((float64((float32(((3.14159 / 1e10) * 3.14159)) * (float32(1.5) + 1.5))) / 1.5) - 1e10) < 0.0))))
    var v4 uint64 = 18446744073709551615
    v11 := GenS1{F0: 2, F1: 65535, F2: ""}
    v12 := []int{128, 1}
    v2 -= v2
    var v13 int32 = -2147483648
    v2 += 628
    v4 += 0
    v14 := v1["foo bar"]
    if (-1 > 64) {
        v14 = v2
        if (!v3) {
            v13 = (v13 | (v13 + (-(v13))))
            _ = len(v12)
            v15 := GenS1{F0: 2, F1: 4294967295, F2: ""}
            _ = v15
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS2{F0: fi16, F1: "\n\t"}
            fmt.Println(v13, v2)
            _ = fi16
            _ = v17
        }
    }
    v3 = v3
    switch "a" {
    case "hello":
        for fi18 := 0; fi18 < 3; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := GenS3{F0: 2}
                _ = fi19
                _ = v20
            }
            if (!v3) {
                var v21 int64 = -1
                for fi22 := 0; fi22 < 19; fi22++ {
                    switch (fi22 >> ((4294967295) & 63)) {
                    case 2147483647:
                        if v3 {
                            var v23 bool = v3
                            if ((0 != 42) || v3) {
                                v24 := GenS1{F0: 42, F1: 18446744073709551615, F2: "\n\t"}
                                if false {
                                    var v25 float32 = 1e10
                                    v1["a"] = (v2 * v2)
                                    _ = v25
                                }
                                if (!false) {
                                    fmt.Println(v23)
                                } else {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    fmt.Println(v3, v14)
                                    v1["hello"] = (^(v2))
                                }
                                _ = v24
                            } else {
                                var v26 int = 841
                                fmt.Println(v3)
                                v14 -= (v2 & v14)
                                _ = v26
                            }
                            _ = v23
                        }
                    }
                    _ = fi22
                }
                fmt.Println(fi18)
                _ = v21
            }
            fmt.Println(v2, v14)
            _ = fi18
        }
    case "a":
        var v27 uint = 42
        _ = v27
    case "\n\t":
        if ((v2 ^ v2) <= uint32(v4)) {
            v28 := GenS4{F0: -128, F1: v4}
            fmt.Println(v3, v2, v13)
            _ = v28
        } else {
            v13 = v13
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            v29 := GenS2{F0: 128, F1: ("\n\t" + (("" + "hello") + "\n\t"))}
            _ = v29
        }
    }
    fmt.Println(v2)
    v13 += (v13 * 902)
    v1[("\n\t" + ("hello" + "foo bar"))] = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
