package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32, p1 int8) int16 {
    switch ("hello" + "a") {
    case "hello":
        var v6 uint8 = 255
        _ = v6
    case "foo bar":
        var v7 int = -9223372036854775808
        _ = v7
    }
    fmt.Println(p1)
    p1 = p1
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = -2.25
    var v3 string = ""
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]int)
    v5[2] = -2147483648
    switch 4294967295 {
    case -1:
        v8 := GenS1{F0: "", F1: 9223372036854775807}
        _ = v8
    case -2147483648:
        v9 := []int{-1, 32768, 9223372036854775807, 2147483647}
        _ = v9
    case 629:
        if false {
            v4[2147483648] = "foo bar"
            fmt.Println(v2, v2)
        }
    default:
        for fi10 := 0; fi10 < 19; fi10++ {
            v2 = ((v2 / 1.5) / -2.25)
            fmt.Println(v2)
            _ = fi10
        }
    }
    var v11 int64 = 42
    v11 = (v11 % ((-397) | 1))
    var v12 bool = false
    v2 = v2
    fmt.Println(v11, v3)
    for fi13 := 0; fi13 < 9; fi13++ {
        v12 = v12
        _ = fi13
    }
    var v14 uint8 = 42
    v1[42] = 42
    switch v3 {
    case "\n\t":
        switch v3 {
        case "":
            v15 := GenS1{F0: v3, F1: -728}
            _ = v15
        case "a":
            if true {
                switch 128 {
                case 128:
                    var v16 uint32 = 291
                    _ = v16
                case 524:
                    if false {
                        switch v3 {
                        case "":
                            for fi17 := 0; fi17 < 5; fi17++ {
                                switch (v3 + ((("hello" + "") + v3) + v3)) {
                                case "a":
                                    v3 = v3
                                case "\n\t":
                                    delete(v5, 4294967295)
                                }
                                fmt.Println(v11)
                                fmt.Println(v14, v14, v12)
                                _ = fi17
                            }
                        case "\n\t":
                            if false {
                                fmt.Println(v2)
                                v18 := GenS1{F0: v3, F1: 127}
                                _ = v18
                            }
                        case "foo bar":
                            var v19 int8 = -1
                            _ = v19
                        default:
                            v20 := []int{661, -844, 32767}
                            _ = v20
                        }
                    }
                default:
                    var v21 float32 = (0.0 - -2.25)
                    _ = v21
                }
            } else {
                switch -1 {
                case 128:
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v11 = (v11 << ((565) & 63))
                        if (0 > 42) {
                            var v23 int32 = -859
                            _ = v23
                        }
                        _ = fi22
                    }
                default:
                    var v24 uint64 = 434
                    _ = v24
                }
                for fi25 := 0; fi25 < 15; fi25++ {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        fmt.Println(v3, v14)
                        switch (-(fi26)) {
                        case -875:
                            if (-2.25 != 1.5) {
                                if v12 {
                                    var v27 bool = v12
                                    _ = v27
                                } else {
                                    var v28 int = (-(fi26))
                                    _ = v28
                                }
                                v2 = v2
                                _ = len(v5)
                            } else {
                                var v29 int32 = 0
                                fmt.Println(v29, fi26, fi26)
                                if v12 {
                                    var v30 int16 = int16(fi25)
                                    _ = v30
                                }
                                v31 := []int{(fi25 - 128), fi26, fi26, 4294967295}
                                _ = v29
                                _ = v31
                            }
                        case -424:
                            v14 -= (^(v14))
                        case 376:
                            v32 := GenS2{F0: 2147483647}
                            _ = v32
                        }
                        v12 = (true || (211 >= 32767))
                        _ = fi26
                    }
                    var v33 int16 = 223
                    _ = fi25
                    _ = v33
                }
            }
        case "hello":
            var v34 uint8 = v14
            _ = v34
        }
    case "a":
        v35 := []int{-424}
        _ = v35
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        v11 -= (v11 + v11)
        for fi37 := 0; fi37 < 15; fi37++ {
            v12 = ((-666 != 0) || (-1 == -524))
            switch v3 {
            case "hello":
                fmt.Println(v3)
            default:
                v12 = (v12 || (1e10 != 0.0))
            }
            v38 := []int{fi37}
            v11 = -295
            _ = fi37
            _ = v38
        }
        var v39 int = fi36
        _ = fi36
        _ = v39
    }
    v3 = v3
    var v40 int8 = 10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
