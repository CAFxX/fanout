package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]string)
    var v3 float64 = 1.5
    var v4 int32 = 1
    var v5 string = "\n\t"
    fmt.Println(v5)
    if (42 <= 704) {
        if (!(!true)) {
            if false {
                if true {
                    var v6 bool = false
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v4)
                    case "":
                        var v7 float64 = v3
                        _ = v7
                    case "foo bar":
                        v8 := GenS1{F0: 143, F1: (v5 + ""), F2: -1}
                        _ = v8
                    }
                    _ = v6
                } else {
                    var v9 string = (v5 + "")
                    var v10 int8 = 0
                    _ = v10
                    _ = v9
                }
                var v11 bool = false
                if false {
                    v12 := []int{}
                    _ = v12
                } else {
                    if false {
                        switch v5 {
                        case "":
                            var v13 int16 = -283
                            _ = v13
                        case "a":
                            if (v11 || (v3 < -2.25)) {
                                if (false && true) {
                                    var v14 int16 = 1
                                    _ = v14
                                } else {
                                    fmt.Println(v1)
                                    v15 := v2[v5]
                                    v5 = v15
                                    _ = v15
                                }
                                if (1.5 != float32((float64(0.1) / 1e10))) {
                                    v16 := GenS2{F0: 571, F1: v1}
                                    _ = v16
                                }
                            }
                        case "foo bar":
                            v17 := GenS3{F0: 42, F1: v4}
                            _ = v17
                        }
                        v18 := []int{2147483647}
                        v11 = ((v1 ^ v1) < v1)
                        _ = v18
                    }
                    v19 := []int{0}
                    var v20 float32 = 0.1
                    _ = v19
                    _ = v20
                }
                _ = v11
            }
            if (true || (v5 < v5)) {
                var v21 uint = 389
                v22 := []int{}
                _ = v21
                _ = v22
            }
        } else {
            if true {
                if ((1.5 * 3.14159) >= float32(0.1)) {
                    var v23 int = int(v4)
                    v24 := GenS2{F0: 1, F1: 65535}
                    _ = v23
                    _ = v24
                }
            } else {
                v25 := v2["\n\t"]
                _ = v25
            }
            v26 := GenS4{F0: 18446744073709551615}
            v27 := GenS5{F0: 88}
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            v3 = float64(0.1)
            v29 := []int{fi28, (fi28 * (fi28 % ((32767) | 1))), fi28}
            _ = fi28
            _ = v29
        }
        v30 := GenS5{F0: 727}
        _ = v30
    } else {
        var v31 float64 = 0.1
        _ = v31
    }
    v5 = v5
    if true {
        v32 := GenS6{F0: false}
        var v33 bool = true
        var v34 uint64 = 451
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 string = "foo bar"
    var v36 uint32 = 65535
    var v37 uint32 = v36
    if false {
        if (18446744073709551615 <= 4294967295) {
            var v38 uint64 = 42
            fmt.Println(v5)
            if false {
                fmt.Println(v1)
                v5 = ""
                var v39 bool = (v3 == v3)
                _ = v39
            } else {
                v5 = ("" + "hello")
                v40 := GenS4{F0: 2}
                _ = v40
            }
            _ = v38
        }
        var v41 int32 = (v4 - 1)
        fmt.Println(v41, v4)
        _ = v41
    }
    var v42 float32 = -2.25
    fmt.Println(v4)
    v35 = "hello"
    fmt.Println(v3, v3, v36)
    var v43 int8 = 1
    var v44 int64 = 1
    if false {
        v45 := GenS2{F0: 65535, F1: v1}
        v46 := v2["hello"]
        _ = v45
        _ = v46
    } else {
        v47 := GenS5{F0: -547}
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
