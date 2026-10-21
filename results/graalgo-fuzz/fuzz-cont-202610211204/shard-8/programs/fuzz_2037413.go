package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 bool; F2 uint16 }
type GenS5 struct { F0 string; F1 uint; F2 int32 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int8; F1 bool; F2 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 1
    var v3 uint16 = 65535
    var v4 uint = 2
    v3 = v3
    v3 = v3
    var v5 float64 = (3.14159 + 1.5)
    v5 -= float64(((0.1 - 1e10) * -2.25))
    switch "hello" {
    case "foo bar":
        v6 := GenS1{F0: 1}
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 42}
        v4 -= v4
        v9 := []int{(^(fi7)), fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        v10 := GenS1{F0: -9223372036854775808}
        v4 = (v4 % (((v4 ^ 4294967295)) | 1))
        _ = v10
    } else {
        v11 := GenS2{F0: 42}
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int32 = (v2 + v2)
            _ = fi12
            _ = v13
        }
        var v14 float32 = -2.25
        _ = v11
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: -9223372036854775808}
        fmt.Println(fi15, v5, fi15)
        var v17 bool = true
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := []int{2147483647, 1, 42}
    for fi19 := 0; fi19 < 2; fi19++ {
        v2 = (v2 << ((v4) & 63))
        var v20 uint32 = 731
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        _ = len(v18)
        _ = fi21
    }
    v22 := GenS1{F0: 0}
    var v23 int32 = (v2 << ((v4) & 63))
    for fi24 := 0; fi24 < 9; fi24++ {
        if true {
            fmt.Println(v4)
            _ = len(v18)
        }
        _ = len(v1)
        for fi25 := 0; fi25 < 8; fi25++ {
            fmt.Println(v3, v23, fi24)
            _ = fi25
        }
        _ = fi24
    }
    fmt.Println(v2, v3)
    switch "a" {
    case "foo bar":
        var v26 int64 = 42
        _ = v26
    case "hello":
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 float64 = v5
            for fi29 := 0; fi29 < 9; fi29++ {
                if false {
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                } else {
                    var v30 uint64 = 106
                    var v31 uint32 = 1
                    var v32 uint32 = v31
                    _ = v30
                    _ = v31
                    _ = v32
                }
                var v33 int = fi27
                var v34 uint32 = 107
                _ = fi29
                _ = v33
                _ = v34
            }
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = fi27
            _ = v28
        }
    case "":
        v35 := GenS1{F0: 42}
        _ = v35
    default:
        var v36 int8 = -128
        _ = v36
    }
    v37 := GenS3{F0: v3}
    if false {
        var v38 float32 = 1e10
        switch "hello" {
        case "hello":
            v39 := GenS4{F0: 42, F1: (!true), F2: v3}
            _ = v39
        case "foo bar":
            if ((42 < 2147483647) || false) {
                v40 := GenS4{F0: -1, F1: (1 >= 280), F2: (v3 % ((v3) | 1))}
                for fi41 := 0; fi41 < 4; fi41++ {
                    switch "hello" {
                    case "\n\t":
                        var v42 int32 = v2
                        _ = v42
                    case "foo bar":
                        v2 = (v23 >> ((18446744073709551615) & 63))
                    case "hello":
                        if false {
                            v43 := GenS5{F0: "hello", F1: v4, F2: v23}
                            v44 := GenS6{F0: 807, F1: 42}
                            _ = v43
                            _ = v44
                        }
                    default:
                        fmt.Println(v3)
                    }
                    _ = fi41
                }
                switch ("a" + "foo bar") {
                case "hello":
                    var v45 int8 = 100
                    _ = v45
                case "":
                    if (true && (false || true)) {
                        if false {
                            fmt.Println(v38)
                        }
                        if false {
                            var v46 uint = v4
                            if (255 > 118) {
                                delete(v1, "foo bar")
                                var v47 int64 = 1
                                v48 := GenS3{F0: 42}
                                _ = v47
                                _ = v48
                            }
                            _ = v46
                        }
                        var v49 uint16 = v3
                        _ = v49
                    }
                }
                _ = v40
            }
        case "":
            var v50 int = 9223372036854775807
            _ = v50
        default:
            switch "\n\t" {
            case "hello":
                if len(v18) > 0 {
                    v18[0] = 4294967295
                }
            }
        }
        for fi51 := 0; fi51 < 13; fi51++ {
            if true {
                var v52 uint32 = 2
                v53 := GenS7{F0: 0, F1: true, F2: 1}
                _ = v52
                _ = v53
            }
            _ = fi51
        }
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
