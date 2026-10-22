package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 42
    var v3 float32 = 0.0
    var v4 uint64 = 4294967295
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]bool)
    switch ("hello" + "\n\t") {
    case "foo bar":
        fmt.Println(v2)
    case "hello":
        v2 = v2
    case "\n\t":
        for fi7 := 0; fi7 < 20; fi7++ {
            switch (fi7 & fi7) {
            case 158:
                for fi8 := 0; fi8 < 15; fi8++ {
                    v4 = v4
                    _ = fi8
                }
            }
            _ = fi7
        }
    }
    v9 := GenS1{F0: 737}
    if ((^(v4)) >= (v4 % (((v4 ^ v4)) | 1))) {
        fmt.Println(v2)
        var v10 string = "a"
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            var v12 int16 = v2
            _ = v12
        }
        var v13 uint64 = 4294967295
        var v14 uint = 561
        _ = fi11
        _ = v13
        _ = v14
    }
    var v15 int8 = -128
    fmt.Println(v15, v15)
    for fi16 := 0; fi16 < 13; fi16++ {
        v4 = v4
        _ = fi16
    }
    var v17 float32 = float32(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
