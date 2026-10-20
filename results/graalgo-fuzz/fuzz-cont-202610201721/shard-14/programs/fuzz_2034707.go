package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint; F2 bool }

func main() {
    v1 := []uint32{4294967295, 65535}
    var v2 uint8 = 225
    v3 := make(map[uint32]bool)
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 127 {
    case -1:
        var v4 int64 = 0
        _ = v4
    case 566:
        if true {
            for fi5 := 0; fi5 < 2; fi5++ {
                for fi6 := 0; fi6 < 11; fi6++ {
                    var v7 string = ("hello" + "")
                    var v8 uint8 = (v2 - (^(v2)))
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                _ = fi5
            }
            v9 := GenS1{F0: 42}
            switch "a" {
            case "a":
                switch ("hello" + "a") {
                case "\n\t":
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                }
            case "":
                fmt.Println(v2)
            case "hello":
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := GenS1{F0: (fi10 << ((1) & 63))}
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    v12 := GenS2{F0: "hello", F1: uint(v2), F2: (!true)}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            _ = v9
        } else {
            fmt.Println(v2)
        }
    case -803:
        var v13 string = ""
        _ = v13
    }
    var v14 int = 965
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 int32 = -1
        v2 += v2
        _ = fi15
        _ = v16
    }
    if true {
        v14 = v14
        var v17 string = "hello"
        _ = v17
    } else {
        v2 -= (v2 >> ((827) & 63))
        fmt.Println(v14, v14)
    }
    v18 := v3[967]
    v18 = false
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 float32 = ((1.5 - float32(((0.1 * (((1.5 + 0.1) + (0.0 * -2.25)) + 3.14159)) - 0.1))) + float32((1.5 - 0.1)))
        var v21 float32 = v20
        switch "a" {
        case "a":
            var v22 int64 = int64(fi19)
            _ = v22
        case "foo bar":
            if v18 {
                v23 := v3[4294967295]
                v21 = ((v20 / float32((1.5 / 1.5))) / float32(3.14159))
                for fi24 := 0; fi24 < 15; fi24++ {
                    v23 = ("" >= ("hello" + "hello"))
                    var v25 uint64 = 1
                    v26 := GenS2{F0: "", F1: 42, F2: v18}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = v23
            } else {
                var v27 float32 = v20
                var v28 bool = (42 >= 32767)
                var v29 uint64 = 489
                _ = v27
                _ = v28
                _ = v29
            }
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    var v30 int32 = 1
    fmt.Println(v18, v18, v14)
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 int64 = 9223372036854775807
        _ = fi31
        _ = v32
    }
    v30 = v30
    v18 = (1 <= 42)
    v14 += (v14 << ((999) & 63))
    var v33 uint32 = 2147483648
    if (!false) {
        v30 = (v30 << ((0) & 63))
        if v18 {
            v34 := v3[65535]
            _ = v34
        } else {
            if (-32768 >= 32767) {
                fmt.Println(v18)
                var v35 bool = v18
                _ = v35
            }
            fmt.Println(v33, v14, v33)
        }
        v36 := GenS1{F0: v14}
        _ = v36
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        if v18 {
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 int32 = v30
                _ = len(v3)
                v40 := v3[(v33 & v33)]
                _ = fi38
                _ = v39
                _ = v40
            }
        }
        v41 := GenS1{F0: -9223372036854775808}
        var v42 uint16 = 0
        _ = fi37
        _ = v41
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v33
}
