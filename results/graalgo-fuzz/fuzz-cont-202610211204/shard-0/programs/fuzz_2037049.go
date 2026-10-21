package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
func genFn0() int8 {
    v6 := GenS1{F0: 0}
    v7 := GenS1{F0: 629}
    if (1 == -115) {
        var v8 uint16 = 65535
        v9 := GenS2{F0: 4294967295}
        v8 = v8
        _ = v8
        _ = v9
    } else {
        var v10 uint16 = 255
        v11 := []int{42, 2147483647, 9223372036854775807, -873}
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = v10
        _ = v11
    }
    var v12 int8 = 0
    _ = v12
    _ = v6
    _ = v7
    return v12
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 42
    var v3 int32 = 0
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    switch 4294967295 {
    case -319:
        var v13 int64 = 1
        _ = v13
    case 357:
        v3 = 757
    }
    var v14 float32 = 1e10
    v5[42] = -9223372036854775808
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int64 = 0
        v4[("" + ("" + ("a" + "")))] = true
        _ = fi15
        _ = v16
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        v17 := GenS1{F0: -9223372036854775808}
        _ = v17
    case "":
        var v18 int16 = 352
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := []int{}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v2 = v2
        v2 += 0
        _ = fi21
    }
    fmt.Println(v14, v14, v2)
    var v22 uint64 = 0
    v23 := v1[1]
    var v24 uint = 18446744073709551615
    v25 := v5[551]
    switch (^(v25)) {
    case -2147483648:
        if true {
            v26 := []int{}
            _ = v26
        }
    default:
        v24 = v2
    }
    v27 := GenS3{F0: -32}
    for fi28 := 0; fi28 < 9; fi28++ {
        if (65535 <= 786) {
            v14 -= 1.5
            fmt.Println(v2)
        } else {
            v24 = (v2 * v2)
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 int32 = 0
                _ = fi29
                _ = v30
            }
            if (515 != 171) {
                v23 -= (-(fi28))
            } else {
                for fi31 := 0; fi31 < 17; fi31++ {
                    fmt.Println(v25, v24)
                    switch 32768 {
                    case -9223372036854775808:
                        var v32 int64 = 0
                        _ = v32
                    default:
                        var v33 uint32 = 967
                        _ = v33
                    }
                    _ = fi31
                }
            }
        }
        switch int(v22) {
        case 2147483647:
            switch "hello" {
            case "\n\t":
                for fi34 := 0; fi34 < 9; fi34++ {
                    v3 = (v3 | v3)
                    var v35 string = "hello"
                    _ = len(v1)
                    _ = fi34
                    _ = v35
                }
            case "a":
                v36 := GenS3{F0: 42}
                _ = v36
            default:
                v37 := GenS4{F0: (336 < 65535)}
                _ = v37
            }
        case -9223372036854775808:
            switch (fi28 + v23) {
            case 127:
                v25 = -618
            case 4294967295:
                switch "hello" {
                case "foo bar":
                    v38 := []int{}
                    _ = v38
                case "\n\t":
                    v39 := []int{(v25 * (fi28 ^ (v25 >> (((v24 | 1)) & 63)))), fi28}
                    _ = v39
                default:
                    if true {
                        _ = len(v1)
                    } else {
                        switch "" {
                        case "a":
                            v3 -= -2147483648
                        }
                        var v40 int8 = -119
                        _ = v40
                    }
                }
            default:
                switch "hello" {
                case "":
                    var v41 string = ""
                    _ = v41
                }
            }
        case -2147483648:
            _ = len(v1)
        }
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
