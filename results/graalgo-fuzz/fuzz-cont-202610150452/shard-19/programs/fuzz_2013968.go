package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 string; F2 uint64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint16) int {
    if false {
        p0 = 255
    } else {
        var v7 float32 = 0.1
        _ = v7
    }
    switch ("foo bar" + "foo bar") {
    case "a":
        v8 := GenS1{F0: 0}
        _ = v8
    }
    v9 := []int{-937, 4294967295, -1, 1}
    _ = p0
    _ = v9
    return 42
}
func genFn1(p0 uint8, p1 uint) bool {
    v10 := []int{1, 42, 112}
    p0 = p0
    p1 = (^(p1))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v10
    return (!true)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (-2147483648 >= -2147483648)
    var v3 bool = (4294967295 >= 858)
    v4 := make(map[uint32]bool)
    var v5 uint = 65535
    v6 := make(map[int]uint32)
    for fi11 := 0; fi11 < 17; fi11++ {
        v4[2147483648] = true
        fmt.Println(v2)
        _ = fi11
    }
    var v12 float32 = -2.25
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            if true {
                var v15 int16 = 42
                fmt.Println(v15, fi14)
                v16 := GenS1{F0: (v15 % (((v15 & (v15 | 624))) | 1))}
                _ = v15
                _ = v16
            } else {
                var v17 uint64 = 605
                _ = v17
            }
            fmt.Println(v12, v3)
            v18 := GenS1{F0: 32767}
            _ = fi14
            _ = v18
        }
        switch fi13 {
        case 128:
            v19 := []int{-547, fi13, fi13, fi13}
            _ = v19
        case -9223372036854775808:
            v20 := GenS1{F0: -735}
            _ = v20
        case 225:
            v21 := []int{fi13}
            _ = v21
        }
        fmt.Println(v5, fi13)
        _ = fi13
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v12 += (v12 / float32(3.14159))
        _ = fi22
    }
    v12 = 3.14159
    for fi23 := 0; fi23 < 7; fi23++ {
        v24 := GenS2{F0: v3, F1: 105, F2: 27}
        _ = fi23
        _ = v24
    }
    fmt.Println(v5)
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 uint64 = 0
        _ = fi25
        _ = v26
    }
    if (-1 < 2147483647) {
        switch ("" + "\n\t") {
        case "foo bar":
            var v27 string = ""
            _ = v27
        case "hello":
            v2 = v2
        case "\n\t":
            v28 := GenS3{F0: 798}
            _ = v28
        default:
            fmt.Println(v5)
        }
    } else {
        v1[1] = 9223372036854775807
        if ((int64(v5) != 42) && true) {
            v29 := GenS4{F0: 9223372036854775807, F1: ("foo bar" + ""), F2: 522}
            if (("hello" + "foo bar") == ("foo bar" + "hello")) {
                if v3 {
                    var v30 int = -641
                    _ = v30
                }
                if ("hello" == "a") {
                    var v31 int8 = 42
                    switch "foo bar" {
                    case "foo bar":
                        for fi32 := 0; fi32 < 11; fi32++ {
                            v33 := GenS3{F0: 18446744073709551615}
                            _ = fi32
                            _ = v33
                        }
                    case "a":
                        v12 += 1.5
                    case "hello":
                        var v34 bool = ((1.5 > (0.0 - (0.1 * 0.1))) || false)
                        _ = v34
                    }
                    _ = v31
                }
            } else {
                var v35 uint = (v5 - v5)
                var v36 bool = (1 != 100)
                _ = v35
                _ = v36
            }
            _ = v29
        } else {
            v37 := v6[-9223372036854775808]
            switch "hello" {
            case "":
                for fi38 := 0; fi38 < 16; fi38++ {
                    switch ("hello" + "\n\t") {
                    case "":
                        var v39 string = "a"
                        _ = v39
                    }
                    _ = fi38
                }
            case "hello":
                _ = len(v4)
            case "foo bar":
                for fi40 := 0; fi40 < 4; fi40++ {
                    v4[v37] = v3
                    for fi41 := 0; fi41 < 9; fi41++ {
                        v12 -= v12
                        switch (("a" + ("hello" + "hello")) + "a") {
                        case "":
                            if (2 == 208) {
                                v5 = (^(v5))
                            }
                        case "foo bar":
                            fmt.Println(v5)
                            fmt.Println(v2)
                            var v42 string = "hello"
                            _ = v42
                        }
                        var v43 float32 = v12
                        _ = fi41
                        _ = v43
                    }
                    _ = fi40
                }
            default:
                delete(v6, 32767)
            }
            if false {
                v44 := GenS2{F0: (1 == 1), F1: 4, F2: 42}
                _ = v44
            } else {
                v5 += (v5 % (((^(v5))) | 1))
                if (true || v2) {
                    fmt.Println(v12, v12, v2)
                    fmt.Println(v37, v5)
                }
                for fi45 := 0; fi45 < 13; fi45++ {
                    v5 = (v5 * 18446744073709551615)
                    v46 := GenS5{F0: 18446744073709551615}
                    fmt.Println(v5, v2)
                    _ = fi45
                    _ = v46
                }
            }
            _ = v37
        }
    }
    v2 = ((v5 | (v5 * 65535)) > v5)
    for fi47 := 0; fi47 < 17; fi47++ {
        v48 := GenS4{F0: -1, F1: "foo bar", F2: 18446744073709551615}
        if v2 {
            v49 := GenS4{F0: 9223372036854775807, F1: "a", F2: 18446744073709551615}
            v2 = v2
            _ = v49
        } else {
            switch "a" {
            case "a":
                v5 = 854
            default:
                v50 := GenS3{F0: 4294967295}
                _ = v50
            }
            v3 = (!(65535 > 512))
            v5 -= 18446744073709551615
        }
        _ = fi47
        _ = v48
    }
    delete(v4, 42)
    v5 = (^(v5))
    v51 := GenS5{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
