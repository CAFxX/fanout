package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    var v1 string = "hello"
    v2 := []string{v1, (v1 + v1)}
    v3 := []int{}
    fmt.Println(v1)
    for fi4 := 0; fi4 < 14; fi4++ {
        if len(v2) > 0 {
            v2[0] = v1
        }
        v1 = v1
        _ = fi4
    }
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 uint32 = 4294967295
        v7 := GenS1{F0: 4294967295, F1: (fi5 >= fi5)}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint8 = 103
    var v9 uint8 = 54
    if true {
        for fi10 := 0; fi10 < 16; fi10++ {
            switch (((v1 + ((v1 + v1) + v1)) + v1) + "\n\t") {
            case "\n\t":
                v9 += (^(v8))
            }
            var v11 uint = 4294967295
            for fi12 := 0; fi12 < 19; fi12++ {
                v8 = (v8 + (v9 & (v9 ^ (v9 | (^(v8))))))
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
    } else {
        var v13 float32 = (0.0 - ((3.14159 + 1e10) * 1.5))
        v8 = 206
        _ = v13
    }
    if (true && (-678 == -796)) {
        var v14 float32 = (float32(1.5) - 1.5)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v15 uint16 = 0
        _ = v14
        _ = v15
    } else {
        v3 = append(v3, 128)
        switch (("a" + "hello") + ((v1 + (v1 + "foo bar")) + v1)) {
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "foo bar":
            switch (v1 + (v1 + (v1 + "\n\t"))) {
            case "a":
                v9 = 255
            case "":
                v1 = v1
            case "hello":
                v8 -= v8
            }
        case "a":
            if false {
                v1 = v1
                if len(v3) > 0 {
                    v3[0] = 32767
                }
                v3 = append(v3, 42)
            }
        default:
            v8 = 217
        }
    }
    switch 127 {
    case -693:
        for fi16 := 0; fi16 < 19; fi16++ {
            v9 = v9
            _ = fi16
        }
    case 4294967295:
        if false {
            v17 := GenS1{F0: 584, F1: false}
            fmt.Println(v9)
            _ = v17
        }
    default:
        v8 -= 127
    }
    var v18 int32 = 0
    v19 := GenS1{F0: 42, F1: false}
    v18 = (v18 | v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
