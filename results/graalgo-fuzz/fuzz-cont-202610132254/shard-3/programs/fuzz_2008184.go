package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (1.5 / 1.0)
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 float64 = 3.14159
        _ = fi4
        _ = v5
    }
    var v6 int64 = 831
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            if (true && false) {
                var v9 int8 = 121
                _ = v9
            }
            _ = fi8
        }
        v10 := GenS1{F0: v6}
        _ = fi7
        _ = v10
    }
    switch ("\n\t" + "a") {
    case "foo bar":
        switch ("foo bar" + ("hello" + "\n\t")) {
        case "foo bar":
            if (-1 >= 42) {
                var v11 int8 = -114
                if false {
                    v2 = v2
                }
                _ = v11
            }
        case "":
            v12 := []int{}
            _ = v12
        case "\n\t":
            delete(v1, "a")
        }
    case "":
        v13 := GenS1{F0: (v6 * v6)}
        _ = v13
    case "a":
        var v14 int64 = (v6 * v6)
        _ = v14
    }
    v6 += (v6 / (((v6 >> ((733) & 63))) | 1))
    switch "foo bar" {
    case "foo bar":
        if (!(!true)) {
            var v15 int = 9223372036854775807
            var v16 uint64 = 4294967295
            for fi17 := 0; fi17 < 7; fi17++ {
                fmt.Println(v15, v2)
                v15 = -1
                _ = fi17
            }
            _ = v15
            _ = v16
        } else {
            var v18 int64 = (v6 & 9223372036854775807)
            var v19 uint16 = 1
            _ = v18
            _ = v19
        }
    case "\n\t":
        v6 = v6
    }
    v20 := v1["\n\t"]
    v20 = ("\n\t" < "foo bar")
    fmt.Println(v6, v6, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v6
}
