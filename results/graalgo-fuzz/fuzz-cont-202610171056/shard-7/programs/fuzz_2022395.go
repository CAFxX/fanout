package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 3.14159
    var v3 uint16 = 2
    v4 := make(map[string]uint32)
    var v5 bool = (261 != 1)
    var v6 uint = 4294967295
    var v7 string = ""
    switch v7 {
    case "":
        switch ("\n\t" + "a") {
        case "hello":
            v8 := []int{}
            _ = v8
        default:
            if (!(v5 || (false || v5))) {
                var v9 int16 = 1000
                v10 := v1[32767]
                if false {
                    var v11 uint16 = 0
                    _ = v11
                } else {
                    _ = len(v1)
                    v3 += (v3 & (v3 ^ v3))
                }
                _ = v10
                _ = v9
            } else {
                var v12 int = -191
                v13 := GenS1{F0: -32768, F1: 0}
                _ = v12
                _ = v13
            }
        }
    case "\n\t":
        fmt.Println(v3, v6, v5)
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := GenS1{F0: -1, F1: 1000}
            var v17 int8 = 100
            var v18 int16 = 755
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        v7 = ((v7 + (("hello" + (v7 + "hello")) + (v7 + v7))) + v7)
        if true {
            var v19 uint32 = 570
            _ = v19
        }
        _ = fi14
    }
    var v20 bool = (127 < 85)
    v6 += v6
    v21 := GenS1{F0: 32767, F1: 388}
    var v22 uint8 = 0
    if (468 >= v6) {
        v1[4294967295] = 42
        if false {
            if (32767 != -144) {
                if false {
                    v7 = (v7 + v7)
                    var v23 int8 = -1
                    v24 := GenS1{F0: 1000, F1: 42}
                    _ = v23
                    _ = v24
                }
            } else {
                switch v7 {
                case "":
                    v25 := v1[1]
                    _ = v25
                case "\n\t":
                    var v26 float32 = (v2 / v2)
                    _ = v26
                case "a":
                    var v27 bool = v5
                    _ = v27
                }
                for fi28 := 0; fi28 < 9; fi28++ {
                    v29 := []int{(^(fi28)), (fi28 & fi28), fi28}
                    _ = fi28
                    _ = v29
                }
            }
        }
    } else {
        if v20 {
            if v5 {
                if v5 {
                    var v30 uint16 = (v3 - v3)
                    for fi31 := 0; fi31 < 17; fi31++ {
                        fmt.Println(v2)
                        _ = fi31
                    }
                    _ = v30
                } else {
                    switch -471 {
                    case -641:
                        delete(v4, "")
                    default:
                        var v32 int32 = -409
                        _ = v32
                    }
                    v3 = uint16(v6)
                }
            } else {
                for fi33 := 0; fi33 < 10; fi33++ {
                    var v34 int = fi33
                    _ = fi33
                    _ = v34
                }
                var v35 float32 = v2
                _ = v35
            }
        } else {
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 int = (-(fi36))
                for fi38 := 0; fi38 < 9; fi38++ {
                    if false {
                        v20 = false
                        var v39 int64 = 0
                        if v20 {
                            var v40 uint16 = 379
                            var v41 uint32 = 0
                            var v42 string = v7
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        _ = v39
                    } else {
                        v22 = v22
                        fmt.Println(v20)
                    }
                    _ = fi38
                }
                v37 = (v37 - fi36)
                _ = fi36
                _ = v37
            }
        }
        switch "a" {
        case "\n\t":
            v22 = (v22 & v22)
        }
    }
    switch 120 {
    case 0:
        fmt.Println(v6, v22, v6)
    case 2147483647:
        for fi43 := 0; fi43 < 11; fi43++ {
            var v44 float64 = 1e10
            _ = fi43
            _ = v44
        }
    case -859:
        v20 = (42 != -10)
    }
    v4["foo bar"] = 2147483647
    switch 187 {
    case 32767:
        var v45 uint8 = (v22 ^ uint8(v3))
        _ = v45
    }
    v5 = ((^(v22)) < 255)
    v6 = (v6 * v6)
    v22 = v22
    if v5 {
        var v46 float64 = (1e10 * 1e10)
        _ = v46
    }
    switch "" {
    case "a":
        var v47 uint = 1
        _ = v47
    }
    v3 = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
