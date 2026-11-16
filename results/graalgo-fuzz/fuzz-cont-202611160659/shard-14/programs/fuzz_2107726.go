package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint) uint32 {
    fmt.Println(p0)
    v7 := []int{}
    var v8 float64 = 3.14159
    var v9 uint = p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{676, 2147483648, 0}
    var v3 uint = 65535
    var v4 uint = v3
    v5 := make(map[int]string)
    var v6 uint32 = 2147483647
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            switch ("hello" + ("a" + "hello")) {
            case "a":
                var v12 string = (("foo bar" + (("" + (("" + "") + ("a" + ""))) + "")) + (("a" + ("\n\t" + "hello")) + "\n\t"))
                _ = v12
            default:
                switch "hello" {
                case "":
                    v13 := GenS1{F0: 165}
                    switch 32768 {
                    case 1:
                        v4 = (v3 >> ((v3) & 63))
                    case 42:
                        v14 := GenS1{F0: -2147483648}
                        _ = v14
                    case -834:
                        v15 := GenS1{F0: 0}
                        _ = v15
                    default:
                        fmt.Println(v4, v4, v3)
                    }
                    switch ((("hello" + "hello") + (("hello" + "foo bar") + "hello")) + "foo bar") {
                    case "foo bar":
                        switch ("\n\t" + ("foo bar" + ("hello" + (("" + "\n\t") + "a")))) {
                        case "":
                            var v16 string = "a"
                            _ = v16
                        case "\n\t":
                            v17 := GenS2{F0: 42}
                            _ = v17
                        }
                    case "a":
                        fmt.Println(v4, v1, fi11)
                    default:
                        v18 := GenS1{F0: -1}
                        _ = v18
                    }
                    _ = v13
                }
            }
            _ = fi11
        }
        fmt.Println(v1, v1, v3)
        fmt.Println(v4, fi10, fi10)
        _ = fi10
    }
    fmt.Println(v3)
    v19 := GenS3{F0: (v6 << (((^(v4))) & 63))}
    if (v1 >= ((v1 - float64(-2.25)) / float64(1e10))) {
        v20 := v5[814]
        fmt.Println(v6, v3)
        _ = v20
    } else {
        var v21 int16 = 32767
        var v22 int8 = -128
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: -1}
    v1 += 0.0
    v6 = v6
    var v24 uint8 = 0
    var v25 uint8 = (v24 + v24)
    v25 = 255
    v26 := v5[-9223372036854775808]
    for fi27 := 0; fi27 < 2; fi27++ {
        fmt.Println(v1, v24, v26)
        for fi28 := 0; fi28 < 15; fi28++ {
            v24 += v24
            _ = fi28
        }
        _ = fi27
    }
    switch -218 {
    case 849:
        var v29 int8 = -24
        _ = v29
    case 127:
        var v30 string = "hello"
        _ = v30
    default:
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 bool = (899 >= 42)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi31
            _ = v32
        }
    }
    v33 := GenS1{F0: 354}
    var v34 int8 = 127
    var v35 string = (("\n\t" + v26) + "hello")
    fmt.Println(v1, v25)
    var v36 string = v26
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
