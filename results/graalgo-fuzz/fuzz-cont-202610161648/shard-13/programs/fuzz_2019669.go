package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int16; F2 string }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[int]int)
    var v2 string = "hello"
    v3 := make(map[string]string)
    var v4 uint = 0
    v2 = v2
    v5 := GenS1{F0: uint16(v4), F1: 679, F2: -2147483648}
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 = 728
        _ = fi6
    }
    v4 = v4
    v2 = "foo bar"
    v7 := GenS2{F0: -2147483648}
    if (true && false) {
        v8 := []int{}
        _ = v8
    } else {
        if true {
            v2 = "a"
            switch 1 {
            case 32768:
                fmt.Println(v2)
            case 127:
                for fi9 := 0; fi9 < 11; fi9++ {
                    if true {
                        var v10 uint = v4
                        _ = v10
                    }
                    switch (v2 + "hello") {
                    case "\n\t":
                        v11 := v1[(fi9 + -9223372036854775808)]
                        _ = v11
                    case "a":
                        if (32767 != 42) {
                            v12 := GenS3{F0: (^(v4)), F1: 42, F2: v2}
                            var v13 int = (fi9 << (((^(v4))) & 63))
                            _ = v12
                            _ = v13
                        } else {
                            for fi14 := 0; fi14 < 15; fi14++ {
                                v15 := []int{(^(fi9)), (fi14 ^ fi9), (^(fi9))}
                                v16 := v1[fi9]
                                var v17 bool = ((!(127 >= 1)) || (-454 < (^(fi9))))
                                _ = fi14
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            var v18 int8 = 16
                            fmt.Println(fi9, fi9)
                            _ = v18
                        }
                    }
                    _ = fi9
                }
            }
            fmt.Println(v2)
        } else {
            if false {
                fmt.Println(v4, v2)
                var v19 int = 9223372036854775807
                v20 := GenS4{F0: true}
                _ = v19
                _ = v20
            }
        }
    }
    v2 = v2
    v3[""] = (v2 + "")
    var v21 uint8 = 128
    v2 = v2
    v22 := GenS1{F0: 2, F1: 255, F2: -639}
    if true {
        v23 := v3[(v2 + (v2 + ("a" + v2)))]
        _ = v23
    } else {
        v24 := GenS4{F0: (v2 < "a")}
        v25 := GenS4{F0: true}
        _ = v24
        _ = v25
    }
    fmt.Println(v2, v21, v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
