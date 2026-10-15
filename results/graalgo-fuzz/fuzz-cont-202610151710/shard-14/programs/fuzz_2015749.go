package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 int8; F1 int32; F2 uint32 }
func genFn0() uint32 {
    v5 := GenS1{F0: 42, F1: 1000}
    switch "\n\t" {
    case "a":
        var v6 int16 = 42
        _ = v6
    }
    var v7 uint16 = 0
    v8 := GenS2{F0: 1, F1: -127}
    _ = v5
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int64, p1 uint8, p2 bool) bool {
    v9 := GenS3{F0: 2}
    if p2 {
        fmt.Println(p2)
        fmt.Println(p2, p1)
    } else {
        var v10 int = -298
        fmt.Println(v10, p2, p1)
        p0 = 42
        _ = v10
    }
    p1 = p1
    if (1 >= -114) {
        var v11 int32 = 2147483647
        fmt.Println(p2)
        _ = v11
    }
    var v12 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return true
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    var v3 uint64 = 1
    var v4 string = ("hello" + "foo bar")
    switch 639 {
    case 9223372036854775807:
        v13 := GenS1{F0: 9223372036854775807, F1: -756}
        _ = v13
    case 127:
        var v14 uint32 = 23
        _ = v14
    case -9223372036854775808:
        var v15 uint16 = 1
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS3{F0: 16}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := v1[v4]
            if (-5 > 127) {
                switch (("" + ((v4 + (v4 + v4)) + v4)) + (v4 + "\n\t")) {
                case "\n\t":
                    var v21 int = (-(fi18))
                    _ = v21
                case "a":
                    var v22 uint32 = 584
                    _ = v22
                case "hello":
                    delete(v1, ("" + "a"))
                }
                v3 = v3
                if (true && v2) {
                    var v23 int = fi19
                    v24 := GenS3{F0: uint8(v3)}
                    fmt.Println(v4, v20, v23)
                    _ = v23
                    _ = v24
                }
                if (!(1 < 65535)) {
                    v25 := GenS1{F0: -9223372036854775808, F1: 1000}
                    var v26 int32 = -35
                    var v27 uint16 = 164
                    _ = v25
                    _ = v26
                    _ = v27
                }
            } else {
                fmt.Println(fi19, v20)
                v28 := []int{120}
                var v29 uint = 699
                _ = v28
                _ = v29
            }
            v30 := []int{}
            v31 := GenS4{F0: (42 <= 2147483647)}
            _ = fi19
            _ = v20
            _ = v30
            _ = v31
        }
        if true {
            fmt.Println(v4)
            var v32 bool = true
            if ("" == v4) {
                if (!(!(!(0.0 > 1e10)))) {
                    v3 = (v3 >> ((509) & 63))
                }
            }
            _ = v32
        } else {
            var v33 string = ((v4 + (v4 + "\n\t")) + v4)
            v33 = v33
            if (!v2) {
                fmt.Println(v33)
            } else {
                v3 = 42
                if (42 != 100) {
                    if (-1 > -764) {
                        if true {
                            for fi34 := 0; fi34 < 17; fi34++ {
                                v35 := GenS5{F0: 892}
                                v36 := GenS1{F0: 1, F1: -1}
                                fmt.Println(v3, v2)
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                        } else {
                            var v37 uint64 = 42
                            _ = v37
                        }
                        v38 := GenS1{F0: 9223372036854775807, F1: -541}
                        var v39 uint8 = 152
                        _ = v38
                        _ = v39
                    } else {
                        var v40 uint32 = 2
                        _ = v40
                    }
                    v41 := []int{fi18, (fi18 >> ((442) & 63)), -9223372036854775808}
                    var v42 uint = 550
                    _ = v41
                    _ = v42
                }
                for fi43 := 0; fi43 < 6; fi43++ {
                    v44 := GenS6{F0: true, F1: 0}
                    v2 = false
                    _ = fi43
                    _ = v44
                }
            }
            _ = v33
        }
        _ = fi18
    }
    var v45 float32 = ((((((-2.25 / 3.14159) / (3.14159 - (1.5 / 0.1))) / -2.25) / 1.5) + 1e10) + float32(1e10))
    v46 := GenS7{F0: 32768, F1: 2147483647}
    v47 := GenS2{F0: 42, F1: -128}
    v3 -= (v3 | 1)
    var v48 int64 = 9223372036854775807
    switch -2147483648 {
    case 413:
        fmt.Println(v48)
    }
    fmt.Println(v48, v2)
    v49 := GenS8{F0: 29, F1: 2147483647, F2: 65535}
    switch -362 {
    case 9223372036854775807:
        var v50 float64 = 1.5
        _ = v50
    default:
        for fi51 := 0; fi51 < 8; fi51++ {
            if (2147483647 == 1) {
                v52 := []int{32768, 127, (fi51 << ((65535) & 63)), fi51}
                _ = v52
            }
            _ = fi51
        }
    }
    _ = len(v1)
    v45 = (0.0 + ((((v45 + float32((1e10 * float64(v45)))) * (v45 + (v45 * (v45 - -2.25)))) / 1.0) + float32(((0.0 + 0.0) - 0.0))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
