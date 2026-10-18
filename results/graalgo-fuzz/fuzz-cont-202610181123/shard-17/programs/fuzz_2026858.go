package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint16; F2 int32 }

func main() {
    var v1 uint = 4294967295
    var v2 uint8 = 42
    v3 := make(map[uint32]int)
    var v4 uint32 = 65535
    var v5 uint = v1
    v3[v4] = 42
    var v6 bool = true
    v7 := GenS1{F0: 42}
    switch 0 {
    case 595:
        var v8 int = -2147483648
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: 0}
        _ = fi9
        _ = v10
    }
    v11 := v3[v4]
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS3{F0: 422}
        v2 = v2
        _ = fi12
        _ = v13
    }
    var v14 float32 = -2.25
    v4 = v4
    var v15 uint8 = 255
    var v16 int8 = 1
    delete(v3, (^(v4)))
    if v6 {
        v1 += 77
    } else {
        var v17 string = ((("\n\t" + "foo bar") + "foo bar") + "foo bar")
        var v18 uint32 = v4
        _ = v17
        _ = v18
    }
    fmt.Println(v15, v6)
    for fi19 := 0; fi19 < 19; fi19++ {
        switch (-(v11)) {
        case 4294967295:
            for fi20 := 0; fi20 < 7; fi20++ {
                v15 -= v15
                _ = fi20
            }
        case 0:
            v11 = v11
        case 662:
            if false {
                switch 9223372036854775807 {
                case 127:
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v22 := GenS2{F0: (v16 << ((v1) & 63))}
                        v16 -= v16
                        _ = fi21
                        _ = v22
                    }
                case -9223372036854775808:
                    var v23 int16 = 1000
                    _ = v23
                }
                v24 := GenS4{F0: v11, F1: 32767, F2: 1}
                _ = v24
            }
        }
        v1 += v5
        var v25 uint64 = 853
        _ = fi19
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 bool = (!v6)
        for fi28 := 0; fi28 < 19; fi28++ {
            v14 = -2.25
            v29 := []int{v11, fi28}
            v27 = false
            v3[(^(v4))] = fi26
            switch "\n\t" {
            case "hello":
                v11 = (fi28 & fi26)
            case "a":
                fmt.Println(v5, v6)
            default:
                var v30 float32 = 0.0
                _ = v30
            }
            var v31 int32 = 42
            var v32 int8 = (v16 * 43)
            _ = fi28
            _ = v29
            _ = v31
            _ = v32
        }
        var v33 int16 = 42
        _ = fi26
        _ = v27
        _ = v33
    }
    v11 += int(v4)
    var v34 uint32 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
