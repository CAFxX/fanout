package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 int8; F2 int16 }
type GenS4 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{42, -2147483648, 2147483647, -173}
    v3 := []uint32{2147483647}
    var v4 uint = 18446744073709551615
    var v5 uint64 = 1
    v6 := make(map[uint32]int)
    var v7 bool = (-909 != 2147483647)
    v8 := GenS1{F0: 42, F1: (32768 > 42), F2: 65535}
    if v7 {
        if false {
            v9 := GenS2{F0: v7, F1: -32768}
            _ = v9
        } else {
            var v10 int8 = 100
            _ = v10
        }
        fmt.Println(v5)
    } else {
        switch -254 {
        case 127:
            var v11 uint64 = v5
            _ = v11
        case 1:
            for fi12 := 0; fi12 < 19; fi12++ {
                fmt.Println(v5, fi12)
                if ((2 == 32767) && (v7 || ((v4 & v4) >= (^(v4))))) {
                    switch fi12 {
                    case 4294967295:
                        fmt.Println(v7)
                    case 1:
                        v3 = append(v3, 65535)
                    case 32767:
                        fmt.Println(v4)
                    }
                }
                _ = fi12
            }
        case 681:
            v13 := GenS2{F0: ((!(!(-2147483648 >= 42))) && (v7 && v7)), F1: 42}
            _ = v13
        }
        if (v7 || v7) {
            var v14 int64 = 1
            switch "foo bar" {
            case "hello":
                var v15 int = 32768
                _ = v15
            case "\n\t":
                switch -2147483648 {
                case 2147483647:
                    switch 1 {
                    case -2147483648:
                        v6[502] = 32767
                    case 169:
                        switch 4294967295 {
                        case 32767:
                            for fi16 := 0; fi16 < 12; fi16++ {
                                if (fi16 >= fi16) {
                                    var v17 int64 = (v14 + (v14 | v14))
                                    _ = v17
                                }
                                _ = len(v1)
                                var v18 uint16 = 0
                                _ = fi16
                                _ = v18
                            }
                        }
                    case 127:
                        v14 = v14
                    default:
                        v5 += 828
                    }
                case -949:
                    v19 := GenS2{F0: true, F1: 845}
                    _ = v19
                case 128:
                    v20 := v6[2]
                    _ = v20
                }
            case "":
                fmt.Println(v4, v14, v4)
            }
            _ = v14
        } else {
            switch "foo bar" {
            case "a":
                if v7 {
                    var v21 bool = ((270 == v5) && (0 == 42))
                    v7 = (1.5 <= 3.14159)
                    _ = v21
                } else {
                    var v22 string = "hello"
                    v22 = v22
                    if (-1 < int64(v5)) {
                        v3 = append(v3, 4294967295)
                        if false {
                            v7 = true
                            _ = len(v2)
                        }
                        v23 := GenS2{F0: v7, F1: 1000}
                        _ = v23
                    } else {
                        v5 -= (v5 << ((995) & 63))
                        v24 := GenS3{F0: -1, F1: 4, F2: 42}
                        var v25 uint = (v4 << ((v4) & 63))
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                }
            case "\n\t":
                var v26 int8 = 1
                _ = v26
            }
        }
    }
    switch "a" {
    case "a":
        v27 := v6[1]
        _ = v27
    }
    fmt.Println(v7)
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 int16 = int16(v5)
        for fi30 := 0; fi30 < 16; fi30++ {
            delete(v6, 42)
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    _ = len(v6)
    v4 = 4294967295
    for fi31 := 0; fi31 < 14; fi31++ {
        if v7 {
            fmt.Println(v7, v7, v4)
        } else {
            v7 = true
            if false {
                v5 += (v5 * v5)
                var v32 int32 = 927
                if false {
                    v5 += (v5 / ((v5) | 1))
                }
                _ = v32
            } else {
                var v33 uint64 = v5
                if false {
                    for fi34 := 0; fi34 < 10; fi34++ {
                        switch "\n\t" {
                        case "a":
                            var v35 uint = v4
                            _ = v35
                        }
                        _ = fi34
                    }
                    var v36 uint64 = (v5 ^ (v33 - 4294967295))
                    _ = v36
                }
                for fi37 := 0; fi37 < 20; fi37++ {
                    v38 := GenS4{F0: 544, F1: 0}
                    v33 -= 18446744073709551615
                    fmt.Println(v5)
                    _ = fi37
                    _ = v38
                }
                _ = v33
            }
        }
        var v39 int32 = 304
        _ = fi31
        _ = v39
    }
    var v40 uint32 = 2
    v41 := GenS3{F0: -962, F1: -128, F2: 355}
    for fi42 := 0; fi42 < 5; fi42++ {
        v43 := GenS3{F0: 0, F1: 0, F2: 0}
        v7 = (false && v7)
        _ = fi42
        _ = v43
    }
    fmt.Println(v5, v5)
    var v44 string = "foo bar"
    if false {
        if v7 {
            v44 = ("a" + v44)
            v45 := v6[(v40 % ((v40) | 1))]
            _ = v45
        } else {
            var v46 int16 = 0
            _ = v46
        }
        var v47 bool = false
        _ = v47
    } else {
        var v48 uint8 = 193
        if v7 {
            var v49 int8 = 0
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1[42] = ((!true) || v7)
            _ = v49
        } else {
            var v50 uint = 42
            _ = v50
        }
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
