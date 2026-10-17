package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 bool }
func genFn0(p0 uint32) uint64 {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{}
        if len(v7) > 0 {
            v7[0] = (fi6 - fi6)
        }
        fmt.Println(p0, p0)
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 44
    switch 42 {
    case 2147483647:
        for fi9 := 0; fi9 < 5; fi9++ {
            if ("" != "foo bar") {
                switch "hello" {
                case "foo bar":
                    fmt.Println(p0, p0)
                }
            } else {
                for fi10 := 0; fi10 < 9; fi10++ {
                    switch "a" {
                    case "foo bar":
                        v8 -= uint8(p0)
                    case "a":
                        var v11 bool = true
                        _ = v11
                    case "hello":
                        var v12 int16 = 32767
                        _ = v12
                    }
                    var v13 uint32 = 2147483647
                    v14 := GenS1{F0: 42, F1: 2147483647}
                    _ = fi10
                    _ = v13
                    _ = v14
                }
                var v15 float32 = 1e10
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 bool = false
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            if (!true) {
                if true {
                    fmt.Println(v8, p0, fi9)
                    v18 := []int{}
                    p0 = p0
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        switch fi19 {
                        case 0:
                            var v20 bool = (169 == 65535)
                            _ = v20
                        }
                        var v21 uint = 65535
                        var v22 int = fi9
                        _ = fi19
                        _ = v21
                        _ = v22
                    }
                    var v23 int16 = 1
                    v23 = v23
                    _ = v23
                }
                if true {
                    v24 := []int{530, 32768, (fi9 >> ((329) & 63)), fi9}
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 int = fi25
                        if ((704 != 94) && (712 >= 255)) {
                            if len(v24) > 0 {
                                v24[0] = 32768
                            }
                        } else {
                            var v27 int = v26
                            _ = v27
                        }
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                }
                for fi28 := 0; fi28 < 1; fi28++ {
                    if true {
                        v29 := []int{fi9}
                        _ = v29
                    } else {
                        v30 := []int{}
                        _ = v30
                    }
                    v31 := GenS1{F0: 320, F1: -1}
                    _ = fi28
                    _ = v31
                }
            }
            v32 := GenS1{F0: 18446744073709551615, F1: -937}
            _ = fi9
            _ = v32
        }
    case -9223372036854775808:
        v8 = v8
    default:
        fmt.Println(p0, v8)
    }
    fmt.Println(v8, v8)
    _ = p0
    _ = v8
    return 452
}

func main() {
    v1 := []uint32{532, 1, 1}
    var v2 bool = ((-812 != 42) && true)
    var v3 int16 = 42
    var v4 int = -2147483648
    v5 := make(map[int]bool)
    var v33 float64 = -2.25
    switch "a" {
    case "\n\t":
        for fi34 := 0; fi34 < 3; fi34++ {
            if v2 {
                switch ("\n\t" + "") {
                case "foo bar":
                    if v2 {
                        fmt.Println(v4, v2)
                        v33 = v33
                    } else {
                        var v35 uint = 185
                        for fi36 := 0; fi36 < 15; fi36++ {
                            var v37 bool = (fi34 <= (fi36 >> ((v35) & 63)))
                            _ = fi36
                            _ = v37
                        }
                        v35 += 65535
                        _ = v35
                    }
                }
            }
            _ = fi34
        }
    }
    fmt.Println(v3, v2)
    var v38 uint32 = 25
    v39 := GenS1{F0: 1, F1: 109}
    v3 = (v3 >> ((408) & 63))
    v40 := GenS2{F0: (^(v38))}
    if (-833 >= 0) {
        v1 = append(v1, 4294967295)
        var v41 int8 = 0
        _ = len(v5)
        _ = v41
    } else {
        v4 = 4294967295
    }
    v2 = (87 == 2)
    if ((v38 ^ (v38 * 648)) >= 412) {
        var v42 int8 = -71
        v5[(^(v4))] = ((!true) && ((^(v42)) == v42))
        _ = v42
    }
    var v43 int32 = 1
    var v44 bool = (v2 && (1 > 1))
    v45 := GenS3{F0: v43, F1: 127, F2: v44}
    v46 := GenS3{F0: (v43 % ((v43) | 1)), F1: (v4 >> ((0) & 63)), F2: v2}
    v43 -= 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
