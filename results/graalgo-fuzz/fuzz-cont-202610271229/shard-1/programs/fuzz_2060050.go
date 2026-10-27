package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    if true {
        var v4 int = 9223372036854775807
        switch v4 {
        case 32767:
            switch "\n\t" {
            case "":
                for fi5 := 0; fi5 < 20; fi5++ {
                    v6 := []int{(v4 * fi5), (v4 & (v4 + fi5)), (-(fi5)), v4}
                    for fi7 := 0; fi7 < 14; fi7++ {
                        if (!(2 < 330)) {
                            v6 = append(v6, v4)
                            _ = len(v2)
                            fmt.Println(fi5, v4, fi5)
                        }
                        var v8 string = (("" + ("a" + "")) + ("" + "a"))
                        _ = fi7
                        _ = v8
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi5
                    _ = v6
                }
            case "hello":
                for fi9 := 0; fi9 < 20; fi9++ {
                    fmt.Println(fi9, v4)
                    var v10 uint8 = 0
                    _ = fi9
                    _ = v10
                }
            case "\n\t":
                var v11 float64 = 0.0
                _ = v11
            }
        case -2147483648:
            if true {
                v3[602] = -796
                var v12 uint = 305
                _ = v12
            } else {
                v4 = v4
            }
        default:
            fmt.Println(v4)
        }
        _ = v4
    } else {
        if true {
            switch "" {
            case "a":
                v13 := GenS1{F0: false, F1: (("a" + ("\n\t" + "hello")) + "foo bar")}
                _ = v13
            case "hello":
                var v14 uint8 = 8
                _ = v14
            }
            var v15 bool = false
            _ = v15
        } else {
            var v16 int = 2147483647
            v16 += v16
            _ = v16
        }
        v17 := []int{635, -9223372036854775808, 9223372036854775807}
        _ = v17
    }
    var v18 uint32 = 854
    var v19 uint16 = 32767
    var v20 int32 = -6
    v21 := GenS1{F0: false, F1: (("foo bar" + ((("foo bar" + "") + ("a" + (("a" + "\n\t") + ("foo bar" + ("a" + "hello"))))) + "a")) + "a")}
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS1{F0: (42 != 42), F1: ("" + "a")}
        var v24 int = fi22
        for fi25 := 0; fi25 < 17; fi25++ {
            fmt.Println(v20, fi22)
            v26 := GenS1{F0: ((1.5 * 1e10) < 1e10), F1: ("\n\t" + "")}
            _ = fi25
            _ = v26
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    var v27 uint16 = 2
    v28 := v3[(v18 & 1)]
    switch 32768 {
    case 4294967295:
        if false {
            if true {
                _ = len(v2)
            } else {
                v29 := v1[v18]
                _ = v29
            }
        } else {
            v30 := GenS2{F0: 2, F1: 1, F2: -9223372036854775808}
            _ = v30
        }
    case -2147483648:
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 int = (fi31 | -405)
            v33 := GenS3{F0: v19, F1: (v19 << ((4294967295) & 63)), F2: 501}
            _ = fi31
            _ = v32
            _ = v33
        }
    default:
        for fi34 := 0; fi34 < 18; fi34++ {
            v35 := GenS3{F0: (v27 % ((v19) | 1)), F1: v27, F2: 1000}
            v27 += 169
            _ = fi34
            _ = v35
        }
    }
    v36 := []int{}
    v37 := v1[(^(v18))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
}
