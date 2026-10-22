package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 string; F2 uint16 }
func genFn0(p0 uint32, p1 int32, p2 int16) int8 {
    var v6 bool = ((42 != 1) || false)
    if (v6 && false) {
        if true {
            p2 = (p2 / ((p2) | 1))
            fmt.Println(p1, v6)
        }
        var v7 int8 = 110
        for fi8 := 0; fi8 < 13; fi8++ {
            v6 = false
            for fi9 := 0; fi9 < 18; fi9++ {
                p2 = (p2 - p2)
                _ = fi9
            }
            _ = fi8
        }
        _ = v7
    } else {
        v10 := []int{int(p2), -275, 127}
        if ((781 < 32767) || v6) {
            var v11 uint64 = 2
            _ = v11
        }
        v6 = (255 < 1)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -2147483648
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    var v5 uint8 = 204
    delete(v1, -2147483648)
    v12 := v4[""]
    fmt.Println(v12, v5, v5)
    v12 = (v12 - v12)
    _ = len(v1)
    if (int32(v12) != (v2 << ((17) & 63))) {
        v5 = (v5 | 12)
        v13 := []int{}
        var v14 uint = 18446744073709551615
        _ = v13
        _ = v14
    } else {
        v15 := GenS1{F0: 766}
        v16 := GenS2{F0: ("a" + "")}
        if true {
            var v17 uint32 = v12
            var v18 int32 = 42
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 19; fi19++ {
                switch "a" {
                case "hello":
                    if false {
                        v20 := GenS2{F0: ("hello" + "hello")}
                        _ = v20
                    } else {
                        if (v12 != v12) {
                            if true {
                                v21 := []int{fi19, fi19, 0}
                                v22 := GenS3{F0: fi19, F1: "foo bar", F2: 0}
                                fmt.Println(v12, v12, fi19)
                                _ = v21
                                _ = v22
                            } else {
                                if true {
                                    if false {
                                        v23 := GenS3{F0: 9223372036854775807, F1: "", F2: 909}
                                        _ = v23
                                    } else {
                                        if true {
                                            v12 += (v12 + v12)
                                            var v24 int32 = -340
                                            _ = v24
                                        }
                                        v12 += v12
                                        v2 += (v2 & v2)
                                    }
                                    var v25 uint = 18446744073709551615
                                    _ = v25
                                } else {
                                    var v26 bool = (false && (!(true && (0 >= 18446744073709551615))))
                                    _ = v26
                                }
                                if true {
                                    if (-32768 != -1) {
                                        switch fi19 {
                                        case 4294967295:
                                            var v27 uint = 2
                                            _ = v27
                                        default:
                                            fmt.Println(v12)
                                        }
                                        for fi28 := 0; fi28 < 1; fi28++ {
                                            if (-2.25 != (3.14159 + 3.14159)) {
                                                if false {
                                                    v29 := v4["hello"]
                                                    v2 += 863
                                                    _ = v29
                                                }
                                            } else {
                                                fmt.Println(fi28, fi28)
                                                v5 = (^(v5))
                                                v30 := GenS2{F0: "\n\t"}
                                                v2 = (^(v2))
                                                _ = v30
                                            }
                                            v31 := GenS1{F0: 4294967295}
                                            v32 := []int{(^(fi19))}
                                            _ = fi28
                                            _ = v31
                                            _ = v32
                                        }
                                        fmt.Println(fi19)
                                    }
                                    v33 := GenS3{F0: -158, F1: "\n\t", F2: 42}
                                    _ = v33
                                }
                            }
                            var v34 int64 = -9223372036854775808
                            _ = v34
                        } else {
                            v5 = 128
                        }
                        switch fi19 {
                        case 128:
                            var v35 string = ("hello" + (("a" + "\n\t") + "\n\t"))
                            _ = v35
                        default:
                            var v36 float64 = (float64(1e10) - float64(0.1))
                            _ = v36
                        }
                    }
                }
                v2 -= (v2 ^ v2)
                v2 += 2147483647
                _ = fi19
            }
            if true {
                if (true && (!true)) {
                    var v37 uint64 = 18446744073709551615
                    _ = v37
                }
                v12 = (^(v12))
            } else {
                v38 := []int{42, int(v2), -760, 9223372036854775807}
                var v39 int32 = v2
                _ = v38
                _ = v39
            }
            fmt.Println(v2, v5, v2)
        }
        _ = v15
        _ = v16
    }
    if false {
        v40 := []int{-1}
        v5 = v5
        _ = v40
    } else {
        v41 := GenS3{F0: 42, F1: "hello", F2: 2}
        _ = v41
    }
    var v42 uint16 = 42
    v43 := GenS1{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
