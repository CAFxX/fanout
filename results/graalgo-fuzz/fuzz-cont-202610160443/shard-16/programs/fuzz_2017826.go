package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint) uint {
    fmt.Println(p0)
    v4 := []int{}
    fmt.Println(p0)
    for fi5 := 0; fi5 < 18; fi5++ {
        p0 += 1
        v4 = append(v4, -758)
        switch ("a" + "") {
        case "hello":
            for fi6 := 0; fi6 < 16; fi6++ {
                fmt.Println(fi5, fi5)
                v7 := GenS1{F0: 0}
                _ = fi6
                _ = v7
            }
        case "a":
            var v8 uint = (p0 / ((p0) | 1))
            _ = v8
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return (p0 & p0)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 3.14159
    var v3 float32 = 0.0
    v1[2147483648] = ("hello" + ("hello" + "foo bar"))
    var v9 uint64 = 995
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if (-2147483648 > fi10) {
                if false {
                    var v12 uint64 = v9
                    _ = len(v1)
                    switch ("\n\t" + "") {
                    case "":
                        v12 = v9
                    default:
                        var v13 int = 401
                        if false {
                            switch (fi10 - (fi11 >> ((genFn0(1)) & 63))) {
                            case -427:
                                var v14 float32 = (v3 * v3)
                                _ = v14
                            case 127:
                                v15 := GenS2{F0: 2, F1: -32768}
                                _ = v15
                            }
                            switch "foo bar" {
                            case "":
                                v16 := GenS3{F0: 0, F1: -9223372036854775808}
                                _ = v16
                            case "hello":
                                var v17 uint32 = 776
                                _ = v17
                            }
                            if true {
                                var v18 bool = (825 > fi10)
                                v19 := GenS4{F0: (0 != 65535)}
                                var v20 float64 = 3.14159
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        }
                        var v21 uint = 905
                        _ = v13
                        _ = v21
                    }
                    _ = v12
                }
            }
            _ = fi11
        }
        v2 += v2
        switch "\n\t" {
        case "a":
            v3 -= v3
        case "\n\t":
            v22 := v1[396]
            _ = v22
        }
        _ = fi10
    }
    v23 := GenS4{F0: true}
    switch 127 {
    case 42:
        v24 := GenS3{F0: 64, F1: -1}
        _ = v24
    case 2147483647:
        if true {
            v3 = v3
            var v25 float64 = v2
            _ = v25
        }
    default:
        if true {
            v26 := GenS1{F0: -32768}
            _ = v26
        } else {
            fmt.Println(v9, v2)
            fmt.Println(v9, v2)
            switch 2147483647 {
            case 127:
                switch (("hello" + ((("hello" + ("foo bar" + ("foo bar" + ""))) + "hello") + ("a" + "\n\t"))) + "a") {
                case "foo bar":
                    fmt.Println(v3)
                case "":
                    for fi27 := 0; fi27 < 18; fi27++ {
                        for fi28 := 0; fi28 < 14; fi28++ {
                            v29 := []int{fi27}
                            var v30 int8 = 100
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        fmt.Println(v9, v9)
                        _ = fi27
                    }
                case "a":
                    switch -1 {
                    case 32768:
                        switch -1 {
                        case 4294967295:
                            for fi31 := 0; fi31 < 13; fi31++ {
                                var v32 uint16 = 149
                                v33 := []int{}
                                var v34 int = (^(fi31))
                                _ = fi31
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                        case 420:
                            var v35 uint = 0
                            _ = v35
                        }
                    case -8:
                        fmt.Println(v3, v3)
                    default:
                        v36 := []int{4294967295, -1, 0, 0}
                        _ = v36
                    }
                }
            case -459:
                v37 := GenS5{F0: 0}
                _ = v37
            case 9223372036854775807:
                var v38 uint16 = 1
                _ = v38
            }
        }
    }
    var v39 uint32 = 692
    fmt.Println(v39, v9, v39)
    _ = len(v1)
    var v40 int16 = 901
    if false {
        v9 = v9
    }
    v2 = 1.5
    fmt.Println(v2)
    v3 -= ((v3 * v3) + v3)
    v1[1] = ("\n\t" + "hello")
    v41 := GenS6{F0: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v9
}
