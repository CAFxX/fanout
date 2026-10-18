package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]string)
    var v3 uint8 = 128
    v4 := []string{"a", "a", (((("hello" + ("\n\t" + "")) + ("" + "hello")) + "\n\t") + ("hello" + ("foo bar" + "a"))), ""}
    switch "\n\t" {
    case "\n\t":
        v3 = (v3 << ((2) & 63))
    default:
        v5 := GenS1{F0: 0}
        _ = v5
    }
    fmt.Println(v3, v1)
    switch "a" {
    case "":
        fmt.Println(v1, v1)
    case "a":
        var v6 uint16 = 2
        _ = v6
    default:
        fmt.Println(v1)
    }
    v7 := GenS1{F0: 813}
    for fi8 := 0; fi8 < 3; fi8++ {
        _ = len(v4)
        _ = fi8
    }
    var v9 string = "a"
    switch 32768 {
    case -1:
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 9; fi11++ {
                v9 = v9
                _ = fi11
            }
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 uint64 = 2
                _ = fi12
                _ = v13
            }
            _ = fi10
        }
    case 32767:
        v3 = (v3 % ((255) | 1))
    case -2147483648:
        fmt.Println(v9, v9, v3)
    }
    var v14 uint = 4294967295
    switch v9 {
    case "\n\t":
        fmt.Println(v3, v3, v1)
    case "hello":
        fmt.Println(v14, v1)
    case "foo bar":
        v1 = 1e10
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        if false {
            fmt.Println(v14)
            var v16 bool = true
            _ = v16
        } else {
            v17 := GenS1{F0: -228}
            _ = v17
        }
        v3 = v3
        _ = fi15
    }
    var v18 int = 684
    v9 = "hello"
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := GenS1{F0: 0}
            v22 := GenS1{F0: 0}
            _ = fi20
            _ = v21
            _ = v22
        }
        if (false || (42 < 642)) {
            switch v9 {
            case "\n\t":
                var v23 int8 = 0
                _ = v23
            case "hello":
                var v24 bool = (false && (-77 != -52))
                _ = v24
            case "foo bar":
                var v25 uint = v14
                _ = v25
            default:
                _ = len(v2)
            }
            fmt.Println(v18)
            var v26 uint16 = 0
            _ = v26
        } else {
            var v27 int16 = 1
            if false {
                switch ("hello" + ("" + "hello")) {
                case "":
                    _ = len(v4)
                case "a":
                    if true {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            fmt.Println(v18, fi28, v1)
                            if (-72 <= -45) {
                                var v29 int8 = -1
                                _ = v29
                            } else {
                                fmt.Println(v14, v14)
                                fmt.Println(v1, fi19)
                                var v30 int8 = 79
                                _ = v30
                            }
                            _ = fi28
                        }
                    } else {
                        for fi31 := 0; fi31 < 10; fi31++ {
                            v1 += float64((0.1 - -2.25))
                            var v32 uint8 = (v3 / ((0) | 1))
                            _ = fi31
                            _ = v32
                        }
                    }
                case "hello":
                    var v33 int32 = 213
                    _ = v33
                }
                var v34 uint = v14
                _ = v34
            } else {
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
            }
            fmt.Println(v27, v3, v9)
            _ = v27
        }
        fmt.Println(v14, fi19)
        _ = fi19
    }
    v9 = v9
    _ = len(v2)
    v35 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v7
    _ = v9
}
