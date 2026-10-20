package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 string; F1 bool; F2 uint }
type GenS6 struct { F0 uint16; F1 int16; F2 int16 }
type GenS7 struct { F0 uint8; F1 int16 }
func genFn0() int64 {
    if true {
        v6 := GenS1{F0: 42, F1: 1}
        v7 := GenS2{F0: 82}
        _ = v6
        _ = v7
    } else {
        var v8 int8 = 0
        _ = v8
    }
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v9 bool = (!((4294967295 > 714) && false))
        _ = v9
    default:
        if false {
            var v10 uint8 = 0
            _ = v10
        } else {
            for fi11 := 0; fi11 < 12; fi11++ {
                if (((fi11 == -570) || (!(42 >= 9223372036854775807))) && (0 != 42)) {
                    for fi12 := 0; fi12 < 18; fi12++ {
                        var v13 int64 = -9223372036854775808
                        if true {
                            if (-32768 < -512) {
                                v14 := []int{(fi12 >> ((121) & 63)), (fi12 | fi11), fi11}
                                var v15 int32 = 792
                                _ = v14
                                _ = v15
                            }
                            var v16 int64 = (-(v13))
                            var v17 int = fi12
                            _ = v16
                            _ = v17
                        } else {
                            switch fi11 {
                            case -109:
                                var v18 uint32 = 65535
                                _ = v18
                            }
                        }
                        _ = fi12
                        _ = v13
                    }
                    for fi19 := 0; fi19 < 3; fi19++ {
                        v20 := []int{(fi11 - fi19), (-(fi19))}
                        v21 := GenS3{F0: 2}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                fmt.Println(fi11)
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 float64 = float64(1.5)
                    v24 := []int{(fi11 % ((fi22) | 1))}
                    fmt.Println(fi22)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = fi11
            }
        }
    }
    return 0
}

func main() {
    var v1 float32 = 1.5
    var v2 int8 = 42
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    v5 := []string{}
    var v25 uint16 = 807
    if false {
        delete(v4, 2147483647)
        v26 := v4[-1]
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := GenS2{F0: 255}
            fmt.Println(v1, v25, v2)
            fmt.Println(v1, v25)
            _ = fi27
            _ = v28
        }
        _ = v26
    } else {
        var v29 string = "foo bar"
        _ = v29
    }
    v25 = v25
    v30 := GenS4{F0: 2147483647, F1: 196}
    _ = len(v5)
    v2 += v2
    switch 127 {
    case 42:
        v31 := GenS5{F0: "foo bar", F1: (-2.25 >= v1), F2: 0}
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            if true {
                v34 := GenS6{F0: 0, F1: 216, F2: 1000}
                _ = v34
            } else {
                v35 := GenS1{F0: -39, F1: 4294967295}
                var v36 bool = (!true)
                var v37 bool = v36
                switch 299 {
                case -1:
                    v38 := GenS7{F0: 144, F1: 627}
                    _ = v38
                }
                var v39 int64 = int64(fi33)
                _ = v35
                _ = v36
                _ = v37
                _ = v39
            }
            var v40 uint = 18446744073709551615
            v5 = append(v5, "foo bar")
            _ = fi33
            _ = v40
        }
        _ = fi32
    }
    fmt.Println(v25, v1)
    delete(v3, "hello")
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
